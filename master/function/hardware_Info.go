package function

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/peer"
	pb "influxdb/grpc"
	"net"
	"os"
	"strings"
	"time"
)

type Server struct {
	pb.UnimplementedSystemMetricsServer
}

type HardwareInfoRecord struct {
	CPUInfo    string   `json:"cpu_info"`
	MemInfo    string   `json:"mem_info"`
	DiskInfo   string   `json:"disk_info"`
	ClientIP   string   `json:"client_ip"`
	ReceivedAt string   `json:"received_at"`
	SystemInfo []string `json:"system_info"`
}

func (s *Server) GethardwareInfo(ctx context.Context, in *pb.HardwareInfo) (*pb.Response, error) {
	// 获取客户端的 IP 地址
	peerInfo, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("无法获取对方IP")
	}
	clientIP := peerInfo.Addr.(*net.TCPAddr).IP.String()

	// 获取当前时间
	receivedAt := time.Now().Format("2006-01-02 15:04:05")

	// 获取硬件信息
	cpuInfoString := fmt.Sprintf("CPU核心数: %d\nCPU名称: %s\nCPU的最大频率: %.2f GHz\n客户端IP: %s\n", in.NumCores, in.ModelName, in.MaxFrequencyGhz, clientIP)
	memInfoString := fmt.Sprintf("内存总量: %.2f GB\n", in.TotalGb)
	diskSizesStr := strings.Join(in.DiskSizeGb, ", ")
	diskInfoString := fmt.Sprintf("磁盘大小: %s\n磁盘块数: %d\n", diskSizesStr, in.NumBlocks)

	// 在服务器端打印系统信息
	fmt.Println(cpuInfoString)
	fmt.Println(memInfoString)
	fmt.Println(diskInfoString)

	// 将记录的值写入 JSON 文件
	record := HardwareInfoRecord{
		CPUInfo:    cpuInfoString,
		MemInfo:    memInfoString,
		DiskInfo:   diskInfoString,
		ClientIP:   clientIP,
		ReceivedAt: receivedAt,
		SystemInfo: in.Systeminfo,
	}

	if err := writeJSONToFile(record, clientIP); err != nil {
		return nil, fmt.Errorf("无法写入JSON文件: %v", err)
	}

	// 返回成功状态码和信息
	return &pb.Response{Response: "200"}, nil
}

func writeJSONToFile(record HardwareInfoRecord, ip string) error {
	filePath := "./data/" + ip + ".json"

	// 删除现有文件
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return err
	}

	// 打开文件以进行写入，如果文件不存在则创建
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将记录编码为 JSON 并写入文件
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(record); err != nil {
		return err
	}

	return nil
}
