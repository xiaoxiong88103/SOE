package montior

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"strings"
)

// 获取物理磁盘数量
func getPhysicalDiskCount() (int, error) {
	ioCounters, err := disk.IOCounters()
	if err != nil {
		return 0, err
	}
	numPhysicalDisks := len(ioCounters)
	return numPhysicalDisks, nil
}

// 获取 CPU 型号名称的字符串表示
func GetCpuModelString(cpuInfo []cpu.InfoStat) string {
	var modelNames []string
	for _, info := range cpuInfo {
		modelNames = append(modelNames, info.ModelName)
	}
	return strings.Join(modelNames, ", ")
}

// 获取 CPU 的最大频率（GHz）
func GetCPUMaxFrequency(cpuInfo []cpu.InfoStat) float64 {
	var maxFreq float64
	for _, info := range cpuInfo {
		freq := float64(info.Mhz) / 1000.0
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	return maxFreq
}

// 获取内存信息并返回
func getMemoryInfo() *mem.VirtualMemoryStat {
	// 获取内存信息
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("无法获取内存信息: %v\n", err)
		return nil
	}
	return memInfo
}
