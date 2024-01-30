package main

import (
	"fmt"
	"influxdb/config"
)

func main() {
	prot, err := config.Dcode_json("web.json", "token_key")
	if err != nil {
		fmt.Println("开启的时候报错:", err)
	}
	fmt.Println(prot)

}
