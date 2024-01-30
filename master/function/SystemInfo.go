package function

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/grpc/peer"
	"influxdb/config"
	pb "influxdb/grpc"
	"net"
	"time"
)

// 服务器接收到客户端系统信息请求时的处理函数
func (s *Server) GetSystemInfo(ctx context.Context, in *pb.SystemInfo) (*pb.Response, error) {
	// 获取客户端的 IP 地址
	peerInfo, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("无法获取对方IP")
	}
	clientIP := peerInfo.Addr.(*net.TCPAddr).IP.String()

	influxdb_link(
		clientIP,
		in.GetBandwidthUsagePerSec(),
		in.GetCpuUsagePercent(),
		in.GetMemoryUsagePercent(),
		in.GetVpuUsagePercent(),
		in.GetNpuUsagePercent(),
		in.GetGpuUsagePercent(),
		in.GetIoReadUsagePercent(),
		in.GetIoWriteUsagePercent(),
		in.GetNetworkUploadUsagePercent(),
		in.GetNetworkDownloadUsagePercent(),
		in.GetNetworkConnections(),
		in.GetSystemLoadAvg(),
		in.GetDiskSizeGbShengyu(),
		in.Time,
	)

	put_time, err := config.DecodeJsonAsInt("client_config.json", "put_time")
	if err != nil {
		put_time := 60
		return &pb.Response{Time: float32(put_time)}, nil
	}
	// 返回成功状态码和信息
	return &pb.Response{Time: float32(put_time)}, nil
}

func influxdb_link(ip string, bandwidth float32, cpu float32, mem float32, vpu float32, npu []float32, gpu float32, ioread float32, iowrite float32, networkup float32, networkload float32, netcon int64, systemaver float32, disksize string, time_node string) {
	// InfluxDB 链接
	client := influxdb2.NewClient(json_plus("url"), json_plus("token"))
	// 获取写入数据的实例
	writeAPI := client.WriteAPI(json_plus("org"), json_plus("databases"))
	// 创建要写入的数据点
	decodeTime, err := time.Parse(time.RFC3339, time_node)
	if err != nil {
		fmt.Println("时间解析出问题了:", err)
	}
	p := influxdb2.NewPointWithMeasurement("system_info").
		AddTag("ip", ip).
		AddField("bandwidth", bandwidth).
		AddField("cpu", cpu).
		AddField("mem", mem).
		AddField("vpu", vpu).
		AddField("npu", npu).
		AddField("gpu", gpu).
		AddField("ioread", ioread).
		AddField("iowrite", iowrite).
		AddField("networkup", networkup).
		AddField("networkload", networkload).
		AddField("netcon", netcon).
		AddField("systemaver", systemaver).
		AddField("disksize", disksize).
		SetTime(decodeTime)
	fmt.Println("解析出来时间是:", decodeTime)
	// 写入数据
	writeAPI.WritePoint(p)

	// 关闭连接
	client.Close()

}
func json_plus(number string) string {
	par, err := config.Dcode_json("config.json", number)
	if err != nil {
		fmt.Println(err)
	}
	return par
}
