package function

import (
	"encoding/json"
	"fmt"
	"influxdb/config"
	pb "influxdb/grpc"
	"io/ioutil"
	"os"
	"path/filepath"
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
	openjson, err_json := config.DecodeJsonAsInt("config", "client_offonline_log_day")
	if err_json != nil {
		return err_json
	}
	// 删除超过xx天的文件
	tenDaysAgo := time.Now().AddDate(0, 0, -+openjson)
	err = deleteOldFiles(fileDir, tenDaysAgo)
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

func deleteOldFiles(dirPath string, threshold time.Time) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if file.ModTime().Before(threshold) {
			err := os.Remove(filepath.Join(dirPath, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
