package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"influxdb/config"
	pb "influxdb/grpc"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	client_offonline pb.SystemMetricsClient
	conn_offonline   *grpc.ClientConn
	once_offonline   sync.Once
	err_offonline    error // 声明一个错误变量，用于在once.Do外部处理错误
)

func main() {
	client, conn, err := link_master_2()
	if err != nil {
		log.Fatalf("连接服务器时出错: %v", err)
	}
	defer conn() // 确保在函数退出时关闭连接

	for {
		// 找到最旧的日志文件
		uri, err := findOldestLogFile()
		if err != nil {
			fmt.Println("查找最旧日志文件时出现错误:", err)
			time.Sleep(10 * time.Second) // 等待一段时间后重试
			continue
		}
		if uri == "" {
			fmt.Println("没有找到日志文件")
			// 如果没有日志文件了则等待指定时间
			clientFileOff, errDecode := config.DecodeJsonAsInt("config.json", "client_file_off_h")
			fmt.Println("没有日志了 这边 按照config的配置休息:", clientFileOff, "小时")
			if errDecode != nil {
				fmt.Println("解析失败:", errDecode)
				time.Sleep(10 * time.Second) // 等待一段时间后重试
				continue
			}
			time.Sleep(time.Duration(clientFileOff) * time.Hour)
			continue
		}
		err = processLogFile(uri, client)
		if err != nil {
			log.Fatalf("处理日志文件时出错: %v", err)
			time.Sleep(10 * time.Second) // 等待一段时间后重试
		}
	}
}

// processLogFile 读取日志文件，发送第一条记录，成功后更新文件
func processLogFile(filePath string, client pb.SystemMetricsClient) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var systemInfos []pb.SystemInfo

	// 按行读取和解析 JSON 数据
	for scanner.Scan() {
		var systemInfo pb.SystemInfo
		err := json.Unmarshal(scanner.Bytes(), &systemInfo)
		if err != nil {
			return fmt.Errorf("JSON 解析错误: %v", err)
		}
		systemInfos = append(systemInfos, systemInfo)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("文件读取错误: %v", err)
	}

	// 检查是否有记录
	if len(systemInfos) == 0 {
		// 如果文件是空的，删除文件
		return os.Remove(filePath)
	}

	// 这里调用您的 GRPC 方法发送第一条记录
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = client.GetSystemInfo(ctx, &systemInfos[0]) // 假设这是发送方法
	if err != nil {
		return err
	}

	// 成功发送后，从切片中移除第一条记录并更新文件
	return updateLogFile(filePath, systemInfos[1:])
}

// findOldestLogFile 在 ./data/ 目录中找到最旧的日志文件名，基于文件名中的日期排序
func findOldestLogFile() (string, error) {
	fileDir := "./data/"
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return "", err
	}

	// 过滤出日志文件
	var logFiles []os.FileInfo
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "_data.log") {
			logFiles = append(logFiles, file)
		}
	}

	// 如果没有日志文件，返回空字符串
	if len(logFiles) == 0 {
		return "", nil
	}

	// 按文件名中的日期排序
	sort.Slice(logFiles, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02_data.log", logFiles[i].Name())
		dateJ, _ := time.Parse("2006-01-02_data.log", logFiles[j].Name())
		return dateI.Before(dateJ)
	})

	// 返回最旧文件的完整路径
	return filepath.Join(fileDir, logFiles[0].Name()), nil
}

func link_master_2() (pb.SystemMetricsClient, func(), error) {
	var maxRetries = 3
	var retryCount int

	for {
		once_offonline.Do(func() {
			// 解析服务器IP
			ip, errDecode := config.Dcode_json("config.json", "serverip")
			if errDecode != nil {
				err_offonline = fmt.Errorf("解析服务器IP出错: %v", errDecode)
				return
			}
			fmt.Println("连接到服务器:", ip)

			// 建立连接
			conn_offonline, err_offonline = grpc.Dial(
				ip,
				grpc.WithInsecure(),
				grpc.WithBlock(),
			)

			if err_offonline != nil {
				fmt.Println("连接失败, 将在一小时后重试")
				return
			}

			// 创建客户端
			client_offonline = pb.NewSystemMetricsClient(conn_offonline)
		})

		if err_offonline == nil {
			// 连接成功
			break
		}

		if retryCount >= maxRetries {
			// 达到最大重试次数
			fmt.Println("达到最大重试次数, 停止尝试")
			return nil, nil, err_offonline
		}

		// 等待一小时后重试
		time.Sleep(1 * time.Hour)
		retryCount++
		once_offonline = sync.Once{} // 重置 once，允许再次尝试
	}

	return client_offonline, func() { conn_offonline.Close() }, nil
}

func updateLogFile(filePath string, systemInfos []pb.SystemInfo) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, si := range systemInfos {
		jsonData, err := json.Marshal(si)
		if err != nil {
			return err
		}
		_, err = file.Write(jsonData)
		if err != nil {
			return err
		}
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}
	return nil
}
