package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"influxdb/config"
	"influxdb/gin_master/function/user/user_login"
	"influxdb/gin_master/routes"
)

func main() {
	route := gin.Default()

	// 注册 swagger 路由
	url := ginSwagger.URL("http://192.168.2.182:8080/docs/swagger.json") // 设置 Swagger JSON 的实际路径
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	route.Static("/docs", "./docs")

	route.POST("/login", user_login.LoginWeb)
	routes.UserRoutes(route) //用户增删改查的

	routes.DatabasesTSDB(route) //查询的

	routes.Basic_func(route) //基础功能的

	routes.Plugins(route) //插件的功能

	prot, err := config.Dcode_json("web.json", "gin_prot")
	if err != nil {
		fmt.Println("开启的时候报错:", err)
	}

	route.Run(":" + prot)
}
