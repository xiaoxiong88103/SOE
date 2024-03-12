package routes

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function/databases"
)

// DatabasesTSDB 数据库相关的路由
func DatabasesTSDB(router *gin.Engine) {
	show := router.Group("/show_db")
	show.POST("/screen", databases.Get_db_screen)      //筛选cpu内存什么的值拿出来
	show.POST("/time", databases.Get_db_time)          //根据时间拿全部的
	show.GET("/latest", databases.GetLatestSystemInfo) //60分钟内最新的数据
}
