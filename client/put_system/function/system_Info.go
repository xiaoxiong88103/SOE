package function

import (
	"context"
	"fmt"
	"influxdb/config"
	pb "influxdb/grpc"
	"log"
	"time"
)

//发送参数的函数
func System_info(client pb.SystemMetricsClient) (float32, error) {
	// 创建一个有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 获取上传下载的值
	network_download, network_upload := network()
	network_all := float32(network_upload + network_download)

	// 获取CPU、内存、IO读写的使用率
	getSystemStats() //初始化IO
	time.Sleep(1 * time.Second)
	ioreading_use, iowrite_use, cpu_use, memory_use := getSystemStats()
	
	// 调用获取GPU使用率的函数
	gpu_use := float32(getGPUUsage())

	time_put_err_out, errDecode := config.DecodeJsonAsInt("config.json", "time_put_err_out")
	time_err := float32(time_put_err_out)
	if errDecode != nil {
		errDecode = fmt.Errorf("解析心跳时间出问题json is time:string(数字) : %v", errDecode)
		return 10, errDecode
	}
	currentTime := time.Now()
	timeStr := currentTime.Format(time.RFC3339)

	systemInfo := pb.SystemInfo{
		BandwidthUsagePerSec:        network_all,
		CpuUsagePercent:             cpu_use,
		MemoryUsagePercent:          memory_use,
		IoReadUsagePercent:          ioreading_use,
		IoWriteUsagePercent:         iowrite_use,
		NetworkUploadUsagePercent:   float32(network_upload),
		NetworkDownloadUsagePercent: float32(network_download),
		GpuUsagePercent:             gpu_use,
		NetworkConnections:          int64(network_get_cont()),
		SystemLoadAvg:               getAvgLoad(),
		DiskSizeGbShengyu:           GetPartitionSpace(),
		VpuUsagePercent:             0,
		NpuUsagePercent:             getNPU(),
		Time:                        timeStr,
	}
	response, err := client.GetSystemInfo(ctx, &systemInfo)
	if err != nil {
		log.Printf("无法发送系统信息: %v", err)
		Set_Global(1)
		return time_err, err
	} else {
		Set_Global(0)
		return response.GetTime(), nil
	}

}
