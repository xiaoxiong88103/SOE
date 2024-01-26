package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"influxdb/config"
	pb "influxdb/grpc"
	"influxdb/master/function"
	"log"
	"net"
	"time"
)

func main() {
	prot, err := config.Dcode_json("config.json", "prot")
	if err != nil {
		fmt.Println(err)
	}

	lis, err := net.Listen("tcp", ":"+prot)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	kasp := keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(getconfig("MaxConnectionIdle")) * time.Second,     // 如果连接空闲30秒，将开始发送心跳
		MaxConnectionAge:      time.Duration(getconfig("MaxConnectionAge")) * time.Minute,      // 连接最大存活时间
		MaxConnectionAgeGrace: time.Duration(getconfig("MaxConnectionAgeGrace")) * time.Second, // 终止前的附加时间
		Time:                  time.Duration(getconfig("Time")) * time.Second,                  // 每30秒检查一次连接的活跃状态
		Timeout:               time.Duration(getconfig("Timeout")) * time.Second,               // 如果心跳响应10秒未收到，则认为连接已断开
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(kasp),
	)

	pb.RegisterSystemMetricsServer(grpcServer, &function.Server{})
	log.Println("服务器在端口" + prot + "上启动")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("服务器运行失败: %v", err)
	}
}

func getconfig(configKey string) int {
	value, err := config.DecodeJsonAsInt("config.json", configKey)
	if err != nil {
		log.Fatalf("获取配置 %s 出错: %v", configKey, err)
	}
	return value
}
