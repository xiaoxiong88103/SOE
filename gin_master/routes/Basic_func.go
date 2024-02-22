package routes

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function/Basic_Functions/Network"
)

// SetupUserRoutes 设置用户相关的路由
func Basic_func(router *gin.Engine) {
	basic := router.Group("/basic")
	//basic.Use(function.AuthMiddleware()) // 在 basic 组中使用 authMiddleware 中间件
	basic.GET("/ping", Network.Ping_node)
	basic.GET("/nmap", Network.Nmap_node)
}
