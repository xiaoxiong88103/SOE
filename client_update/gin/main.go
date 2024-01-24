package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"influxdb/config"
)

func main() {

	route := gin.Default()

	route.GET("/ping", Get_Null)
	route.GET("/")

	prot, err := config.Dcode_json("gin.json", "prot")
	if err != nil {
		fmt.Println("开启的时候报错:", err)
	}

	route.Run(":" + prot)
}
