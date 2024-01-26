package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"influxdb/client_update/gin/function"
	"influxdb/config"
)

func main() {
	route := gin.Default()
	route.GET("/ping", function.Get_Null)
	update := route.Group("/update")
	update.POST("/aii")
	prot, err := config.Dcode_json("config.json", "gin_prot")
	if err != nil {
		fmt.Println("开启的时候报错:", err)
	}

	route.Run(":" + prot)
}
