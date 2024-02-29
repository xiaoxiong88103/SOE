package main

import (
	"fmt"
	"influxdb/config"
)

func main() {
	dir := config.Gin_master_dir_user()
	fmt.Println(dir)
}
