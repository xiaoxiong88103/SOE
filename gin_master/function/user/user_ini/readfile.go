package user_ini

import (
	"github.com/go-ini/ini"
	"influxdb/config"
)

var Path_ini_file = config.Path_user_ini

// 读取ini文件的 key 和 value功能
func readSectionValues(sectionName string) (map[string]string, error) {
	// 打开配置文件
	cfg, err := ini.Load(Path_ini_file + "user.ini")
	if err != nil {
		return nil, err
	}

	// 获取指定部分
	section := cfg.Section(sectionName)

	// 读取部分中的所有键值对
	sectionValues := make(map[string]string)
	for _, key := range section.KeyStrings() {
		sectionValues[key] = section.Key(key).String()
	}

	return sectionValues, nil
}

// 接收到[XX]对应的KEY 返回value
func getValueInSection(sectionName, keyName string) (string, error) {
	// 读取部分的所有键值对
	sectionValues, err := readSectionValues(sectionName)
	if err != nil {
		return "", err
	}

	// 查找键对应的值
	value, ok := sectionValues[keyName]
	if !ok {
		return "", nil
	}

	return value, nil
}
