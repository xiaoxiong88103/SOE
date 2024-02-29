package main

import (
	"fmt"
	"influxdb/client/ApiGin"
	"influxdb/client/put_system/function"
	"influxdb/client/put_system/link"
	"influxdb/config"
	"log"
	"time"
)

func main() {
	function.Set_Global(1)
	//	开启离线监控日志
	go func() {
		//默认等待10秒来防止出问题
		time.Sleep(10 * time.Second)
		fmt.Println("开启成功单机日志监控联机检查")
		for {
			time_put_err_out, errDecode := config.DecodeJsonAsInt("config.json", "time_put_err_out")
			time.Sleep(time.Duration(time_put_err_out) * time.Second)
			if errDecode != nil {
				fmt.Println("获取出问题了时间:", errDecode)
				time.Sleep(time.Duration(time_put_err_out) * time.Second)
			}
			open := function.Show_Global()
			if open == 1 {
				data := function.Json_data() // 占位符，实际获取数据的方法
				if err := function.Off_line(data); err != nil {
					fmt.Printf("记录数据时出错: %v\n", err)
				}

			}
		}
	}()
	//  开启同步功能
	go func() {
		function.Put_log_json()
	}()

	go func() {
		ApiGin.ApiGin() //开启API的功能给master调用
	}()

	// 使用 Link_master 函数连接服务器
	client, conn, err := link.Link_master()
	if err != nil {
		log.Fatalf("连接服务器时出错: %v", err)
	}
	defer conn() // 确保在函数退出时关闭连接

	// 发送系统参数信息
	function.SendhardwareInfo(client)
	function.Set_Global(0)
	for {
		timeSleep, _ := function.System_info(client)
		//fmt.Println("ok", timeSleep)
		time.Sleep(time.Duration(timeSleep) * time.Second)
	}

}
