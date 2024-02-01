package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

// 初始化为全局变量以跟踪上一次读写的字节数
var prevReadBytes, prevWriteBytes uint64

//获取CPU、内存、IO读写的使用率
func getSystemStats() (float32, float32, float32, float32) {
	// 获取IO读写使用率
	ioStats, err := disk.IOCounters()
	if err != nil {
		fmt.Printf("Error getting IO counters: %v\n", err)
		return 0, 0, 0, 0
	}

	var totalReadBytes, totalWriteBytes uint64
	for _, stat := range ioStats {
		totalReadBytes += stat.ReadBytes
		totalWriteBytes += stat.WriteBytes
	}

	// 计算读写字节增量
	readBytes := totalReadBytes - prevReadBytes
	writeBytes := totalWriteBytes - prevWriteBytes

	// 更新上一次的读写字节总数
	prevReadBytes = totalReadBytes
	prevWriteBytes = totalWriteBytes

	// 将增量转换为MB/s
	ioReadUsage := float32(readBytes) / 1024 / 1024
	ioWriteUsage := float32(writeBytes) / 1024 / 1024

	// 获取CPU使用率
	cpuStats, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Printf("Error getting CPU percent: %v\n", err)
		return 0, 0, 0, 0
	}
	cpuUsage := float32(cpuStats[0])

	// 获取内存使用率
	memStats, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting virtual memory: %v\n", err)
		return 0, 0, 0, 0
	}
	memUsage := float32(memStats.UsedPercent)

	return ioReadUsage, ioWriteUsage, cpuUsage, memUsage
}

func main() {
	// 第一次调用前获取初始IO统计数据
	getSystemStats()
	// 等待1秒
	time.Sleep(1 * time.Second)
	// 获取并打印1秒内的使用量
	ioreading_use, iowrite_use, cpu, mem := getSystemStats()
	fmt.Printf("IO读使用了: %.2f MB\n", ioreading_use)
	fmt.Printf("IO写使用了: %.2f MB\n", iowrite_use)
	fmt.Printf("cpu写使用了: %.2f MB\n", cpu)
	fmt.Printf("mem写使用了: %.2f MB\n", mem)

}
