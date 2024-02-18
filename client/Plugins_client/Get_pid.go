package Plugins_client

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"time"
)

// Get_AIIServr_PID 示例，获取AIIServer0的PID
func Get_AIIServr_PID(processName string) []int32 {
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

	memPercent, err := p.MemoryPercent()
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

	return memPercent, cpuPercent, runTime, nil
}
