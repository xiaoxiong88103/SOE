package main

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	//token := "gE76J6WGq8ze_rcsPdHz1ujZ7eGfUE-x09nK6ueGoInI_59bJW15l4NLaE7tgtN62A7yP-hpLcaAVkZE5tqQYA=="
	token_master := "wu4gt4dawUhisRvG7F_C0yoHVoRBPr3xZnf6lVQQ2o_Dh7gJU6tsGglViR6C1GYmkhft9YgchEnEIC-eanazVw=="
	bucket := "master"
	url := "http://192.168.200.134:8086"
	client := influxdb2.NewClient(url, token_master)
	org := "md"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	p := influxdb2.NewPointWithMeasurement("cpu_usage").
		AddTag("servername", "server1").
		AddField("usage", 10).
		SetTime(time.Now())

	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}
}
