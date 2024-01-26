package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

// 使用指针存储上一次的读写字节总数，便于检查是否已经设置
var prevReadBytes, prevWriteBytes *uint64

// 获取CPU、内存、IO读写的使用率
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

	var ioReadUsage, ioWriteUsage float32

	// 如果是第一次调用，仅设置初始值
	if prevReadBytes == nil || prevWriteBytes == nil {
		prevReadBytes = new(uint64)
		prevWriteBytes = new(uint64)
		*prevReadBytes = totalReadBytes
		*prevWriteBytes = totalWriteBytes
	} else {
		// 计算读写字节增量
		readBytes := totalReadBytes - *prevReadBytes
		writeBytes := totalWriteBytes - *prevWriteBytes

		// 将增量转换为MB/s
		ioReadUsage = float32(readBytes) / 1024 / 1024
		ioWriteUsage = float32(writeBytes) / 1024 / 1024

		// 更新上一次的读写字节总数
		*prevReadBytes = totalReadBytes
		*prevWriteBytes = totalWriteBytes
	}

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
	for {
		read, write, cpu, mem := getSystemStats()
		fmt.Printf("IO Read: %.2f MB/s, IO Write: %.2f MB/s, CPU: %.2f%%, MEM: %.2f%%\n", read, write, cpu, mem)
		time.Sleep(1 * time.Second)
	}
}
