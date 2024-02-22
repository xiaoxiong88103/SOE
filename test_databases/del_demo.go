package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	// 连接到InfluxDB
	const url = "http://192.168.2.182:8086"
	const org = "md"
	const token = "fibRIvanZtNnlQQ33XdVQRKd7tSsbApyvYy4IHIT6IbWb_bL24Lx2ddYvvA2ySzay9oYYaYGcUWLpSN04_OpBQ=="
	client := influxdb2.NewClient(url, token)
	defer client.Close()
	const bucket = "master"

	// 设置要删除数据的时间范围
	start := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC) // Unix 纪元开始时间
	stop := time.Now().Add(24 * time.Hour)               // 当前时间 + 24小时（作为一个未来的时间点）

	// 获取删除API
	deleteAPI := client.DeleteAPI()

	// 执行删除操作
	err := deleteAPI.DeleteWithName(context.Background(), org, bucket, start, stop, "")
	if err != nil {
		fmt.Printf("删除数据时发生错误: %v\n", err)
		return
	}

	fmt.Println("删除操作成功执行")
}
