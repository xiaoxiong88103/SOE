package function

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"influxdb/client/montior"
	pb "influxdb/grpc"
	"log"
	"os/exec"
	"strings"
	"time"
)

func SendhardwareInfo(client pb.SystemMetricsClient) error {
	// 创建一个有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 获取CPU信息
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatalf("无法获取CPU信息: %v", err)
	}

	// 调用 linux_system.go 中的函数获取 CPU 型号和最大频率
	modelName := montior.GetCpuModelString(cpuInfo)
	maxFrequencyGhz := montior.GetCPUMaxFrequency(cpuInfo)

	// 内存的
	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("无法获取内存信息:", err)
	}

	// 内存总量（以 GB 为单位）
	totalGB := float64(memory.Total) / 1024 / 1024 / 1024

	// 获取磁盘大小信息和磁盘数量
	disksize, disknumber := disk_get()

	// 发送 gRPC 请求并传递所有信息
	response, err := client.GethardwareInfo(ctx, &pb.HardwareInfo{
		NumCores:        int32(len(cpuInfo)),
		ModelName:       modelName,
		MaxFrequencyGhz: float32(maxFrequencyGhz),
		TotalGb:         float32(totalGB),
		DiskSizeGb:      disksize,
		NumBlocks:       int32(disknumber),
	})

	if err != nil {
		log.Printf("无法发送CPU信息: %v", err)
		return err
	} else {
		fmt.Println("服务端返回的信息:", response.GetResponse())
	}
	return nil
}

func disk_get() ([]string, int) {
	var disks []string
	var totalDisks int // 记录磁盘数量的变量

	// 执行 fdisk -l 命令获取磁盘信息
	out, err := exec.Command("fdisk", "-l").Output()
	if err != nil {
		fmt.Println("执行 fdisk 命令出错:", err)
		return disks, 0
	}

	// 将命令输出按行分割
	lines := strings.Split(string(out), "\n")

	// 筛选并收集磁盘大小信息
	for _, line := range lines {
		if strings.HasPrefix(line, "Disk") && strings.Contains(line, ":") {
			parts := strings.Split(line, ",")
			if len(parts) > 1 {
				diskInfo := strings.TrimSpace(parts[0])
				disks = append(disks, diskInfo)
				totalDisks++
			}
		}
	}

	return disks, totalDisks
}
