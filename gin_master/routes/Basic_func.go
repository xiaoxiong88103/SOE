package routes

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function/Basic_Functions/Network"
	"influxdb/gin_master/function/Basic_Functions/basic"
)

// SetupUserRoutes 设置用户相关的路由
func Basic_func(router *gin.Engine) {

	//这里是文件管理的代码
	files := router.Group("/files")
	{
		// 设置静态文件服务器，允许浏览和下载文件
		files.GET("/files/list", basic.GetFilelist)
		// 上传文件的接口
		files.PUT("/files/upload", basic.UploadFile)
		//删除函数
		files.DELETE("/files/:filename", basic.RemoveFile)
	}

	basic := router.Group("/basic")
	//basic.Use(function.AuthMiddleware()) // 在 basic 组中使用 authMiddleware 中间件
	basic.GET("/ping", Network.Ping_node)
	basic.GET("/nmap", Network.Nmap_node)
}
