package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"influxdb/client_update/gin/function"
	"influxdb/client_update/gin/function/Plugins"
	"influxdb/config"
)

func main() {
	route := gin.Default()
	route.GET("/ping", function.Get_Null)
	//MVP更新部分
	mvp := route.Group("/mvp")
	mvp.POST("/wget", Plugins.WgetMVP)
	mvp.GET("/os", Plugins.MVPOS)
	mvp.GET("/version", Plugins.GetVersionTxt)

	prot, err := config.Dcode_json("config.json", "gin_prot")
	if err != nil {
		fmt.Println("开启的时候报错:", err)
	}

	route.Run(":" + prot)
}
