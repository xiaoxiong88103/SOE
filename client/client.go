package main

import (
	"fmt"
	"influxdb/client/put_system/function"
	"influxdb/client/put_system/link"
	"log"
	"time"
)

//func main() {
//	// 使用 connectToMetricsServer 函数连接服务器
//	client, conn, err := link.Link_master()
//	if err != nil {
//		log.Fatalf("连接服务器时出错: %v", err)
//	}
//	defer conn() // 确保在函数退出时关闭连接
//
//	// 发送系统参数信息
//	function.SendhardwareInfo(client)
//
//}

func main() {
	// 使用 connectToMetricsServer 函数连接服务器
	client, conn, err := link.Link_master()
	if err != nil {
		log.Fatalf("连接服务器时出错: %v", err)
	}
	defer conn() // 确保在函数退出时关闭连接

	// 发送系统参数信息
	function.SendhardwareInfo(client)

	for {
		timeSleep, _ := function.System_info(client)
		fmt.Println("ok", timeSleep)
		time.Sleep(time.Duration(timeSleep) * time.Second)
	}
}
