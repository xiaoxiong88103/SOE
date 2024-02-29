package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var Path_user_ini, Path_files, Path_config, Path_data, Path_basic string

func Dirfile(dirname string) string {
	// 获取当前执行文件的完整路径
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return "nil"
	}

	// 获取执行文件所在的目录
	execDir := filepath.Dir(execPath)

	// 构建相对于执行文件位置的路径
	relativePath := filepath.Join(execDir, dirname)
	addxiegang := relativePath + "/"
	return addxiegang
}

// 全局变量做文件目录获取用的
func init() {
	Path_user_ini = Dirfile("./data_web/user/")
	Path_files = Dirfile("./files/")
	Path_config = Dirfile("./config/")
	Path_data = Dirfile("./data/")
	Path_basic = Dirfile("./data_web/basic/")
}
