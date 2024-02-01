package routes

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function/plugins/basic"
	mvp2 "influxdb/gin_master/function/plugins/mvp"
	"net/http"
)

// Plugins 插件相关的路由
func Plugins(router *gin.Engine) {
	// 设置静态文件服务器，允许浏览和下载文件
	router.StaticFS("/files", http.Dir("./files"))
	// 上传文件的接口
	router.PUT("/files/upload", basic.UploadFile)
	//删除函数
	router.DELETE("/files/:filename", basic.RemoveFile)

	//mvp插件开发
	mvp := router.Group("/mvp")
	mvp.POST("/version", mvp2.Version_mvp) //查询xx ip机器到版本mvp的
	mvp.POST("/update", mvp2.Update_mvp)   //文件更新到xx机器里
}
