package function

import (
	"github.com/go-ini/ini"
	"io/ioutil"
)

var Path_ini_file = "./data_web/user/"

//读取token文件的功能
func ReadTokenFromFile(filePath string) (string, error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//读取ini文件的 key 和 value功能
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

//接收到[XX]对应的KEY 返回value
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
