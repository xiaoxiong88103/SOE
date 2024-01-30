package routes

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function/databases"
)

// SetupUserRoutes 设置用户相关的路由
func DatabasesTSDB(router *gin.Engine) {
	show := router.Group("/show_db")
	show.POST("/screen", databases.Get_db_screen) //筛选cpu内存什么的值拿出来
	show.POST("/time")                            //根据时间拿全部的
	show.POST("/all")                             //当前30分内的所有都拿出来
}
