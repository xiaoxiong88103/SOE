package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type QueryResult struct {
	Time  string
	Field string
	Value string
}

func main() {
	// 连接到 InfluxDB
	const url = "http://192.168.2.182:8086"
	const org = "md"
	const token = "fibRIvanZtNnlQQ33XdVQRKd7tSsbApyvYy4IHIT6IbWb_bL24Lx2ddYvvA2ySzay9oYYaYGcUWLpSN04_OpBQ=="
	client := influxdb2.NewClient(url, token)
	defer client.Close()
	const bucket = "master"
	// 创建用于存储所有查询结果的切片
	var results []QueryResult

	// 执行第一次查询
	timeRangeStart := "-2h"  // 过去1小时
	timeRangeStop := "now()" // 直到当前时间

	// 定义查询
	query1 := fmt.Sprintf(`
from(bucket: "%s")
  |> range(start: %s, stop: %s)
  |> filter(fn: (r) => r["_measurement"] == "system_info")
  |> filter(fn: (r) => r["_field"] == "cpu"  or r["_field"] == "gpu" or r["_field"] == "ioread" or r["_field"] == "iowrite" or r["_field"] == "netcon" or r["_field"] == "mem" or r["_field"] == "networkload" or r["_field"] == "networkup" or r["_field"] == "systemaver"  or r["_field"] == "vpu" or r["_field"] == "bandwidth")
  |> yield(name: "mean")`, bucket, timeRangeStart, timeRangeStop)

	result, err := client.QueryAPI(org).Query(context.Background(), query1)
	if err != nil {
		fmt.Printf("第一次查询出错: %v\n", err)
		return
	}

	// 解析并存储第一次查询结果
	for result.Next() {
		record := result.Record()
		results = append(results, QueryResult{
			Time:  record.Time().Format(time.RFC3339),
			Field: record.Field(),
			Value: fmt.Sprintf("%v", record.Value()),
		})
	}

	if result.Err() != nil {
		fmt.Printf("在处理第一次查询结果时发生错误: %v\n", result.Err())
	}

	// 执行第二次查询
	query2 := `from(bucket: "master")
		|> range(start: -10m)
		|> filter(fn: (r) => r._measurement == "system_info" and (r._field == "npu" or r._field == "disksize"))`

	result, err = client.QueryAPI(org).Query(context.Background(), query2)
	if err != nil {
		fmt.Printf("第二次查询出错: %v\n", err)
		return
	}

	// 解析并存储第二次查询结果
	for result.Next() {
		record := result.Record()
		results = append(results, QueryResult{
			Time:  record.Time().Format(time.RFC3339),
			Field: record.Field(),
			Value: fmt.Sprintf("%v", record.Value()),
		})
	}

	if result.Err() != nil {
		fmt.Printf("在处理第二次查询结果时发生错误: %v\n", result.Err())
	}

	// 打印所有结果
	for _, res := range results {
		fmt.Printf("时间: %s, 字段: %s, 值: %s\n", res.Time, res.Field, res.Value)
	}
}
