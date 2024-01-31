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
		|> filter(fn: (r) => r._measurement == "system_info" and (r._field == "npu" or r._field == "disksize"))`

	// 执行查询
	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		fmt.Printf("查询出错: %v\n", err)
		return
	}

	// 处理查询结果
	for result.Next() {
		record := result.Record()
		field := record.Field()
		if field == "npu" || field == "disksize" {
			// 转换时间为本地时区
			localTime := record.Time().Local()

			// 输出结果，这里简单地将值作为interface{}类型打印
			fmt.Printf("Time: %s, %s: %v\n", localTime.String(), field, record.Value())
		}
	}

	// 检查查询过程中是否有错误发生
	if result.Err() != nil {
		fmt.Printf("处理查询结果时发生错误: %v\n", result.Err().Error())
	}
}
