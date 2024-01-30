package user_ini

import (
	"fmt"
	"github.com/go-ini/ini"
)

// validateUser 验证用户名和密码是否匹配
func ValidateUser(username, password string) bool {
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
func AddUserdata(username string, password string, admin string, group string) error {
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
	adminSection.NewKey(username, admin) // 100 是管理员的默认值，你可以根据需要修改

	groupSection := cfg.Section("group")
	groupSection.NewKey(username, group)
	// 保存配置文件
	err = cfg.SaveTo(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	return nil
}

// RemoveUserdata 从ini文件中删除用户数据
func RemoveUserdata(username string) error {
	// 打开配置文件
	cfg, err := ini.Load(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	// 删除user部分中的用户数据
	userSection := cfg.Section("user")
	userSection.DeleteKey(username)

	// 删除admin部分中的用户数据
	adminSection := cfg.Section("admin")
	adminSection.DeleteKey(username)

	// 删除group部分中的用户数据
	groupSection := cfg.Section("group")
	groupSection.DeleteKey(username)

	// 保存配置文件
	err = cfg.SaveTo(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserdata 根据用户名更新用户数据
func UpdateUserdata(username string, newPassword string, newAdmin string, newGroup string) error {
	// 打开配置文件
	cfg, err := ini.Load(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	// 获取并更新user部分
	userSection := cfg.Section("user")
	if key, err := userSection.GetKey(username); err == nil {
		key.SetValue(newPassword)
	} else {
		return err // 用户名不存在
	}

	// 获取并更新admin部分
	adminSection := cfg.Section("admin")
	if key, err := adminSection.GetKey(username); err == nil {
		key.SetValue(newAdmin)
	} else {
		return err // 用户名不存在
	}

	// 获取并更新group部分
	groupSection := cfg.Section("group")
	if key, err := groupSection.GetKey(username); err == nil {
		key.SetValue(newGroup)
	} else {
		return err // 用户名不存在
	}

	// 保存配置文件
	err = cfg.SaveTo(Path_ini_file + "user.ini")
	if err != nil {
		return err
	}

	return nil
}
