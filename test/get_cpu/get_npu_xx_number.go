package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	// 连接到InfluxDB
	const url = "http://192.168.2.182:8086"
	const org = "md"
	const token = "fibRIvanZtNnlQQ33XdVQRKd7tSsbApyvYy4IHIT6IbWb_bL24Lx2ddYvvA2ySzay9oYYaYGcUWLpSN04_OpBQ=="
	client := influxdb2.NewClient(url, token)
	defer client.Close()
	const bucket = "master"

	// 创建查询
	query := `from(bucket: "master")
		|> range(start: -10m)
		|> filter(fn: (r) => r._measurement == "system_info" and r._field == "npu")`

	// 执行查询
	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		fmt.Printf("查询出错: %v\n", err)
		return
	}

	// 处理查询结果
	for result.Next() {
		if result.Record().Field() == "npu" {
			// 转换时间为本地时区
			localTime := result.Record().Time().Local()

			// 此处假设NPU值可以被正常读取
			// 由于无法直接处理[]float32类型，您可能需要根据实际情况进行调整
			npuValue := result.Record().Value()
			fmt.Printf("Time: %s, NPU: %v\n", localTime.String(), npuValue)
		}
	}

	// 检查查询过程中是否有错误发生
	if result.Err() != nil {
		fmt.Printf("处理查询结果时发生错误: %v\n", result.Err().Error())
	}
}
