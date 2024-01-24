package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func Dcode_json(configjson string ,Select string) (string, error) {
	// 打开配置文件
	configFile, err := os.Open("config/"+configjson)
	if err != nil {
		return "", fmt.Errorf("无法打开配置文件: %v", err)
	}
	defer configFile.Close()

	// 创建一个map用于存储解析后的JSON数据
	var config map[string]string

	// 使用json.Decoder解码JSON数据
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return "", fmt.Errorf("解码JSON失败: %v", err)
	}

	// 获取serverip的值
	serverIP, found := config[Select]
	if !found {
		return "", fmt.Errorf("未找到" + Select + "配置")
	}

	return serverIP, nil
}

func DecodeJsonAsInt(configjson string,Select string) (int, error) {
	// 打开配置文件
	configFile, err := os.Open("config/"+configjson)
	if err != nil {
		return 0, fmt.Errorf("无法打开配置文件: %v", err)
	}
	defer configFile.Close()

	// 创建一个map用于存储解析后的JSON数据
	var config map[string]string

	// 使用json.Decoder解码JSON数据
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return 0, fmt.Errorf("解码JSON失败: %v", err)
	}

	// 获取指定键的值
	value, found := config[Select]
	if !found {
		return 0, fmt.Errorf("未找到 %s 配置", Select)
	}

	// 将字符串转换为整数
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("配置值 %s 转换为整数时出错: %v", value, err)
	}

	return intValue, nil
}
