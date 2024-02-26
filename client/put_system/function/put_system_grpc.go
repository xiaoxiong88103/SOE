package function

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"influxdb/config"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// GPU的函数调用方法-------------------------------------------------------
// queryGPUInfo 用于查询GPU的特定信息
func queryGPUInfo(query string) ([]string, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu="+query, "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return strings.Split(strings.TrimSpace(string(output)), "\n"), nil
}

// formatGPUInfo 用于根据条件格式化GPU信息
func formatGPUInfo(temperatures, memoryUsed, memoryTotal, utilization []string) ([]string, error) {
	var formattedInfo []string
	for i := 0; i < len(temperatures); i++ {
		infoParts := []string{fmt.Sprintf("GPU%d:", i)}
		infoParts = append(infoParts, fmt.Sprintf("Temp:%s°C", temperatures[i]))
		infoParts = append(infoParts, fmt.Sprintf("Used:%sMB", memoryUsed[i]))
		infoParts = append(infoParts, fmt.Sprintf("Total:%sMB", memoryTotal[i]))
		infoParts = append(infoParts, fmt.Sprintf("Utilization:%s%%", utilization[i]))
		formattedInfo = append(formattedInfo, strings.Join(infoParts, ", "))
	}

	return formattedInfo, nil
}

//GPU的函数调用方法-------------------------------------------------------

// 获取网络上传下载的当前速率的
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

// IO当前读写的
var prevReadBytes, prevWriteBytes uint64

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

// 获取GPU的使用率的
func getGPUUsage() []string {
	arch := runtime.GOARCH
	if arch != "amd64" {
		filePath := "/sys/devices/platform/fb000000.gpu/devfreq/fb000000.gpu/load"
		// 读取文件内容
		content, err := os.ReadFile(filePath)
		if err != nil {
			// 读取文件失败时，返回错误信息的格式化字符串
			return []string{fmt.Sprintf("[Error: %v]", err)}
		}

		// 将文件内容转换为字符串
		loadInfo := string(content)

		// 使用字符串分割函数获取@前面的值
		parts := strings.Split(loadInfo, "@")
		if len(parts) != 2 {
			// 文件内容格式不符合预期时，返回一个错误信息
			return []string{"[Error: parsing GPU load]"}
		}

		// 提取@前面的值并返回
		loadValue := strings.TrimSpace(parts[0])

		return []string{fmt.Sprintf("[GPU:%s]", loadValue)}
	}
	gpu, err := config.Dcode_json("config.json", "gpu")
	if err != nil {
		return []string{fmt.Sprintf("[Error gpu_config: %v]", err)}
	}

	if gpu == "open" {
		// 分别查询GPU温度、已用内存、总内存和利用率
		temperatures, err := queryGPUInfo("temperature.gpu")
		if err != nil {
			return []string{fmt.Sprintf("[Error: %v]", err)}
		}

		memoryUsed, err := queryGPUInfo("memory.used")
		if err != nil {
			return []string{fmt.Sprintf("[Error: %v]", err)}
		}

		memoryTotal, err := queryGPUInfo("memory.total")
		if err != nil {
			return []string{fmt.Sprintf("[Error: %v]", err)}
		}

		utilization, err := queryGPUInfo("utilization.gpu")
		if err != nil {
			return []string{fmt.Sprintf("[Error: %v]", err)}
		}

		// 格式化并输出GPU信息，根据控制变量决定显示哪些信息
		formattedInfo, err := formatGPUInfo(temperatures, memoryUsed, memoryTotal, utilization)
		if err != nil {
			return []string{fmt.Sprintf("[Error: %v]", err)}
		}

		return formattedInfo
	}
	return []string{fmt.Sprintf("")}
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

// getNPU 函数用于从指定文件中获取 NPU 的数值，并返回一个包含三个数值的切片。
func getNPU() []float32 {
	arch := runtime.GOARCH
	if arch == "amd64" {
		// 如果架构是 amd64，返回三个零值的切片
		return []float32{0, 0, 0}
	}

	// 读取文件内容
	content, err := os.ReadFile("/sys/kernel/debug/rknpu/load")
	if err != nil {
		// 如果无法读取文件，返回三个数值都为0的切片，并打印错误信息
		fmt.Println("Error reading file:", err)
		return []float32{0, 0, 0}
	}

	// 将文件内容转换为字符串，并移除换行符
	loadData := strings.ReplaceAll(string(content), "\n", "")

	// 解析字符串中的数值
	var core0, core1, core2 float32
	_, err = fmt.Sscanf(loadData, "NPU load:  Core0: %f%%, Core1: %f%%, Core2: %f%%", &core0, &core1, &core2)
	if err != nil {
		// 如果无法解析字符串，返回三个数值都为0的切片，并打印错误信息
		fmt.Println("Error parsing load data:", err)
		return []float32{0, 0, 0}
	}

	// 返回解析得到的数值组装成的切片
	return []float32{core0, core1, core2}
}

// 获取网络连接数
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
