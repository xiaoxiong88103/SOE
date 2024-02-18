package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

// Get_AIIServr_PID 示例，获取AIIServer0的PID
func Get_AIIServr_PID() []int32 {
	processName := "AIIServer0"
	pids, err := findPIDByName(processName)
	if err != nil {
		fmt.Printf("Error finding process: %s\n", err)
		return nil
	}
	if len(pids) == 0 {
		fmt.Printf("No process found with name %s\n", processName)
		return nil
	}
	return pids
}

// findPIDByName 查找具有指定名称的进程PID
func findPIDByName(processName string) ([]int32, error) {
	var pids []int32
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}
		if name == processName {
			pids = append(pids, p.Pid)
		}
	}

	return pids, nil
}

// GetProcessInfo 获取特定PID进程的内存、CPU使用率和运行时间
func GetProcessInfo(pid int32) (float32, float64, time.Duration, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return 0, 0, 0, err
	}

	memPercent, err := p.MemoryInfo()
	if err != nil {
		return 0, 0, 0, err
	}

	cpuPercent, err := p.CPUPercent()
	if err != nil {
		return 0, 0, 0, err
	}

	createTime, err := p.CreateTime()
	if err != nil {
		return 0, 0, 0, err
	}

	// 计算运行时间
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	runTime := time.Duration(currentTime-int64(createTime)) * time.Millisecond

	return float32(memPercent.RSS), cpuPercent, runTime, nil
}

// readVersionFromFile 读取并解析版本信息
func readVersionFromFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// 使用正则表达式匹配版本信息
	re := regexp.MustCompile(`Version: ([\d\.]+)`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		// 未找到版本信息
		return "", fmt.Errorf("version information not found")
	}

	// 返回匹配的版本号
	return matches[1], nil
}

// getHandleCount 获取指定PID进程的句柄数（文件描述符数量）
func getHandleCount(pid int32) (int, error) {
	fdPath := filepath.Join("/proc", fmt.Sprintf("%d", pid), "fd")
	files, err := ioutil.ReadDir(fdPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 进程可能已经结束
			return 0, fmt.Errorf("process with PID %d does not exist", pid)
		}
		return 0, err
	}
	return len(files), nil
}

func main() {
	// 获取AIIServer0的PID（这里简化，假设你已经有这个函数）
	pids := Get_AIIServr_PID() // 假设这个函数存在，返回AIIServer0的PID列表
	if len(pids) == 0 {
		fmt.Println("AIIServer0 process not found.")
		return
	}
	pid := pids[0] // 取第一个PID

	// 获取进程信息
	memPercent, cpuPercent, runTime, err := GetProcessInfo(pid)
	if err != nil {
		fmt.Printf("Error getting process info: %s\n", err)
		return
	}

	// 获取句柄数
	handleCount, err := getHandleCount(pid)
	if err != nil {
		fmt.Printf("Error getting handle count: %s\n", err)
		return
	}

	// 获取内存信息
	vmem, _ := mem.VirtualMemory()

	// 获取系统负载
	avg, _ := load.Avg()

	// 获取系统启动时间
	upTime, _ := host.Uptime()
	upSince := time.Now().Add(-time.Second * time.Duration(upTime))

	// 读取版本信息
	version, err := readVersionFromFile("/opt/AII/GPU-0/version.txt")
	if err != nil {
		fmt.Printf("Error reading version file: %s\n", err)
		return
	}
	mempercent := memPercent / 1024 / 1024
	// 打印信息
	fmt.Printf("%s | aiiver: %12s | pid:%8d | fd:%5d | cpu:%6.1f | mem:%7.2f | free:%7d | loadavg:%5.2f | uptime: %s | aiirun:%14s |.\n",
		upSince.Format("2006-01-02 15:04:05"), version, pid, handleCount, cpuPercent, mempercent, vmem.Used/1024/1024, avg.Load1, upSince.Format("2006-01-02 15:04:05"), runTime)
}
