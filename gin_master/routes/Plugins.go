package routes

import (
	"github.com/gin-gonic/gin"
	mvp2 "influxdb/gin_master/function/plugins/mvp"
)

// Plugins 插件相关的路由
func Plugins(router *gin.Engine) {
	//mvp插件开发
	mvp := router.Group("/mvp")
	mvp.POST("/version", mvp2.Version_mvp) //查询xx ip机器到版本mvp的
	mvp.POST("/update", mvp2.Update_mvp)   //文件更新到xx机器里
}
