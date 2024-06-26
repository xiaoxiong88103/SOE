package bf_ini

import (
	"fmt"
	"github.com/go-ini/ini"
	"influxdb/config"
	"strconv"
)

// GetINIValue 从INI文件中获取指定section和key的值
func Get_config(section, key string) (string, error) {
	cfg, err := ini.Load(config.Path_basic + "basic.ini")
	if err != nil {
		return "", err
	}

	sec := cfg.Section(section)
	if sec == nil {
		return "", fmt.Errorf("Section not found: %s", section)
	}

	value := sec.Key(key).String()
	return value, nil
}

func Get_config_int(ssection, skey string) (int, error) {
	valuestring, err := Get_config(ssection, skey)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(valuestring)
	if err != nil {
		return 0, fmt.Errorf("转换失败:", err)
	}

	return value, nil
}
