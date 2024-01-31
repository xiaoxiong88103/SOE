package routes

import (
	"github.com/gin-gonic/gin"
)

// Plugins 插件相关的路由
func Plugins(router *gin.Engine) {
	mvp := router.Group("/mvp")
	mvp.PUT("/upload")  //上传文件
	mvp.GET("/files")   //查询文件
	mvp.GET("/version") //查询xx ip机器到版本mvp的
	mvp.POST("/update") //文件更新到xx机器里
}
