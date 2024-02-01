package function

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

//获取网络上传下载的当前速率的
func network() (float64, float64) {
	lastStat, _ := net.IOCounters(true)
	time.Sleep(1 * time.Second)
	newStat, _ := net.IOCounters(true)

	// 计算每秒的上传和下载速率
	for i := range newStat {
		recvDiff := float64(newStat[i].BytesRecv - lastStat[i].BytesRecv)
		transmitDiff := float64(newStat[i].BytesSent - lastStat[i].BytesSent)
		// 将速率从字节转换为MB
		uploadRate := recvDiff / 1024 / 1024
		downloadRate := transmitDiff / 1024 / 1024
		return uploadRate, downloadRate
	}

	return 0, 0
}

// 获取CPU、内存、IO读写的使用率
//func getSystemStats() (float32, float32, float32, float32) {
//	var prevReadBytes, prevWriteBytes uint64
//	// 获取IO读写使用率
//	ioStats, err := disk.IOCounters()
//	if err != nil {
//		fmt.Printf("Error getting IO counters: %v\n", err)
//		return 0, 0, 0, 0
//	}
//
//	var totalReadBytes, totalWriteBytes uint64
//	for _, stat := range ioStats {
//		totalReadBytes += stat.ReadBytes
//		totalWriteBytes += stat.WriteBytes
//	}
//
//	// 计算读写字节增量
//	readBytes := totalReadBytes - prevReadBytes
//	writeBytes := totalWriteBytes - prevWriteBytes
//
//	// 更新上一次的读写字节总数
//	prevReadBytes = totalReadBytes
//	prevWriteBytes = totalWriteBytes
//
//	// 将增量转换为MB/s
//	ioReadUsage := float32(readBytes) / 1024 / 1024
//	ioWriteUsage := float32(writeBytes) / 1024 / 1024
//
//	// 获取CPU使用率
//	cpuStats, err := cpu.Percent(0, false)
//	if err != nil {
//		fmt.Printf("Error getting CPU percent: %v\n", err)
//		return 0, 0, 0, 0
//	}
//	cpuUsage := float32(cpuStats[0])
//
//	// 获取内存使用率
//	memStats, err := mem.VirtualMemory()
//	if err != nil {
//		fmt.Printf("Error getting virtual memory: %v\n", err)
//		return 0, 0, 0, 0
//	}
//	memUsage := float32(memStats.UsedPercent)
//
//	return ioReadUsage, ioWriteUsage, cpuUsage, memUsage
//}

//IO当前读写的
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

//获取GPU的使用率的
func getGPUUsage() float32 {
	cmd := exec.Command("nvidia-smi", "--query-gpu=utilization.gpu", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return 0.0
	}

	gpuUsageStr := strings.TrimSpace(string(output))
	gpuUsage := 0.0
	// 尝试将获取到的GPU使用率转换为float值
	_, err = fmt.Sscanf(gpuUsageStr, "%f", &gpuUsage)
	if err != nil {
		return 0.0
	}

	return float32(gpuUsage)
}

// 系统平均负载
func getAvgLoad() float32 {
	var totalLoad float64
	for i := 0; i < 3; i++ {
		avg, err := load.Avg()
		if err != nil {
			return 0
		}
		totalLoad += avg.Load1
	}
	return float32(totalLoad / 10)
}

//获取NPU的数值
func getNPU() []float32 {
	content, err := ioutil.ReadFile("/sys/kernel/debug/rknpu/load")
	if err != nil {
		// 如果无法读取文件，返回三个数值都为0的切片
		return []float32{0, 0, 0}
	}

	// 将文件内容转换为字符串，并移除换行符
	loadData := strings.ReplaceAll(string(content), "\n", "")

	// 解析字符串中的数值
	var core0, core1, core2 float32
	fmt.Sscanf(loadData, "NPU load:  Core0: %f%%, Core1: %f%%, Core2: %f%%", &core0, &core1, &core2)

	// 将解析得到的数值组装成切片并返回
	return []float32{core0, core1, core2}
}

//获取网络连接数
func network_get_cont() int {
	connections, err := net.Connections("all")
	if err != nil {
		log.Printf("Error getting network connections: %v", err)
		return 0
	}

	return len(connections)
}

// GetPartitionSpace 返回描述所有分区剩余空间的字符串（以 GB 为单位）
func GetPartitionSpace() string {
	var result string

	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatalf("Error retrieving partitions: %v", err)
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			log.Printf("Error getting usage for partition %s: %v", p.Mountpoint, err)
			continue
		}

		remainingGB := float64(usage.Free) / 1024 / 1024 / 1024
		partitionInfo := fmt.Sprintf("Partition: %s, Remaining space: %.2f GB\n", p.Mountpoint, remainingGB)
		result += partitionInfo
	}

	return result
}
