package link

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"influxdb/config"
	pb "influxdb/grpc"
	"sync"
	"time"
)

var (
	client pb.SystemMetricsClient
	conn   *grpc.ClientConn
	once   sync.Once
	err    error // 声明一个错误变量，用于在once.Do外部处理错误
)

func Link_master() (pb.SystemMetricsClient, func(), error) {
	once.Do(func() {
		// 解析服务器IP
		ip, errDecode := config.Dcode_json("config.json", "serverip")
		if errDecode != nil {
			err = fmt.Errorf("解析服务器IP出错: %v", errDecode)
			return
		}
		fmt.Println("连接到服务器:", ip)

		time_put, errDecode := config.DecodeJsonAsInt("config.json", "time_put")
		if errDecode != nil {
			err = fmt.Errorf("解析心跳时间出问题json is time:string(数字) : %v", errDecode)
			return
		}
		fmt.Println("发送心跳每:", time_put, "秒发送一次")

		time_out, errDecode := config.DecodeJsonAsInt("config.json", "time_out")
		if errDecode != nil {
			err = fmt.Errorf("解析心跳时间出问题json is time:string(数字) : %v", errDecode)
			return
		}
		fmt.Println("心跳响应超时时间:", time_out)

		// 设置心跳参数
		kacp := keepalive.ClientParameters{
			Time:                time.Duration(time_put) * time.Second, // 每30秒发送一次心跳
			Timeout:             time.Duration(time_out) * time.Second, // 心跳响应超时时间
			PermitWithoutStream: true,                                  // 即使没有活跃的流也允许发送心跳
		}

		// 建立连接
		conn, err = grpc.Dial(
			ip,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(kacp), // 添加心跳参数
		)

		if err != nil {
			return
		}

		// 创建客户端
		client = pb.NewSystemMetricsClient(conn)
	})

	// 如果once.Do中有错误发生，这里将返回错误
	if err != nil {
		return nil, nil, err
	}

	// 返回客户端和关闭连接的函数
	return client, func() { conn.Close() }, nil
}
