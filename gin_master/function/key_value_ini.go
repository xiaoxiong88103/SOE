package function

import (
	"fmt"
	"github.com/go-ini/ini"
)

// validateUser 验证用户名和密码是否匹配
func validateUser(username, password string) bool {
	// 读取用户名对应的密码
	userPassword, err := getValueInSection("user", username)
	if err != nil || userPassword == "" {
		return false
	}

	// 比较提供的密码与配置中的密码
	return userPassword == password
}

// validateUser 验证用户名和密码是否匹配
func Show_admin(username string) string {
	// 读取用户名对应的密码
	username, err := getValueInSection("user", username)
	if err != nil || username != "" {
		return fmt.Sprintf("错误了用户名是空的:", username)
	}

	// 比较提供的密码与配置中的密码
	return username
}

// 添加用户数据到ini文件
func addUserdata(username string, password string, admin string) error {
	// 打开配置文件
	cfg, err := ini.Load(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	// 获取user部分
	userSection := cfg.Section("user")

	// 添加用户数据
	userSection.NewKey(username, password)
	
	adminSection := cfg.Section("admin")
	adminSection.NewKey(username, "100") // 100 是管理员的默认值，你可以根据需要修改

	// 保存配置文件
	err = cfg.SaveTo(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	return nil
}
