package function

import (
	"encoding/json"
	"fmt"
	"influxdb/config"
	pb "influxdb/grpc"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Off_line 将 &systemInfo 转换为 JSON 并存储在 './data/日期_data.log' 中
func Off_line(systemInfo *pb.SystemInfo) error {
	// 将 systemInfo 转换为 JSON
	jsonData, err := json.Marshal(systemInfo)
	if err != nil {
		return err
	}

	// 生成当前日期的文件名
	fileDir := "./data/"
	fileName := fmt.Sprintf("%s%s_data.log", fileDir, time.Now().Format("2006-01-02"))

	// 确保文件夹存在
	err = os.MkdirAll(fileDir, 0755)
	if err != nil {
		return err
	}
	openjson, err_json := config.DecodeJsonAsInt("config.json", "client_offonline_log_day")
	if err_json != nil {
		return err_json
	}
	// 删除超过xx天的文件
	err = deleteOldFiles(fileDir, openjson)
	if err != nil {
		return err
	}

	// 打开文件以追加数据，如果文件不存在则创建
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将 JSON 数据追加到文件
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	// 可选：追加换行符以分隔每条记录
	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}

func deleteOldFiles(dirPath string, days int) error {
	threshold := time.Now().AddDate(0, 0, -days) // 设置天数阈值

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	fmt.Println("检查文件夹:", dirPath)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// 提取文件名中的日期部分并进行格式检查
		splitName := strings.Split(file.Name(), "_")
		if len(splitName) != 2 || !strings.HasSuffix(splitName[1], ".log") {
			continue
		}

		fileDate, err := time.Parse("2006-01-02", splitName[0])
		if err != nil {
			fmt.Println("日期格式错误:", file.Name())
			continue
		}

		// 比较日期
		if fileDate.Before(threshold) {
			filePath := filepath.Join(dirPath, file.Name())
			fmt.Println("删除过期文件:", filePath, "文件日期:", fileDate)
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
