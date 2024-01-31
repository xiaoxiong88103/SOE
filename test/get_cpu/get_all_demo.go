package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	// 连接到 InfluxDB
	const url = "http://192.168.2.182:8086"
	const org = "md"
	const token = "fibRIvanZtNnlQQ33XdVQRKd7tSsbApyvYy4IHIT6IbWb_bL24Lx2ddYvvA2ySzay9oYYaYGcUWLpSN04_OpBQ=="
	client := influxdb2.NewClient(url, token)
	defer client.Close()

	// 设置查询参数
	const bucket = "master"
	timeRangeStart := "-2h"  // 过去1小时
	timeRangeStop := "now()" // 直到当前时间

	// 定义查询
	query := fmt.Sprintf(`
from(bucket: "%s")
  |> range(start: %s, stop: %s)
  |> filter(fn: (r) => r["_measurement"] == "system_info")
  |> filter(fn: (r) => r["_field"] == "cpu"  or r["_field"] == "gpu" or r["_field"] == "ioread" or r["_field"] == "iowrite" or r["_field"] == "netcon" or r["_field"] == "mem" or r["_field"] == "networkload" or r["_field"] == "networkup" or r["_field"] == "systemaver"  or r["_field"] == "vpu" or r["_field"] == "bandwidth")
  |> yield(name: "mean")`, bucket, timeRangeStart, timeRangeStop)

	// 执行查询
	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		fmt.Printf("查询出错: %v\n", err)
		return
	}

	// 解析并打印结果
	for result.Next() {
		record := result.Record()
		// 将值转换为字符串
		valueStr := fmt.Sprintf("%v", record.Value())
		fmt.Printf("时间: %s, 字段: %s, 值: %s\n", record.Time().Format(time.RFC3339), record.Field(), valueStr)
	}

	if result.Err() != nil {
		fmt.Printf("在获取结果时出现错误: %v\n", result.Err())
	}
}
