package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"influxdb/config"
	pb "influxdb/grpc"
	"influxdb/master/function"
	"log"
	"net"
	"time"
)

// TokenInterceptor 是一个 gRPC 拦截器，用于验证客户端提供的 token
func tokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "无法获取 metadata")
	}

	// 获取客户端提供的 token
	token := md.Get("token")
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "未提供 token")
	}

	// 验证 token 是否等于 "xiaoxiong"
	if token[0] != "xiaoxiong" {
		return nil, status.Errorf(codes.PermissionDenied, "无效的 token")
	}

	// 如果 token 有效，则调用下一个处理程序
	return handler(ctx, req)
}

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
		grpc.UnaryInterceptor(tokenInterceptor), // 添加拦截器
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
