package routes

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function"
	"influxdb/gin_master/function/user_login"
)

// SetupUserRoutes 设置用户相关的路由
func UserRoutes(router *gin.Engine) {
	user := router.Group("/user_login")
	user.Use(function.AuthMiddleware()) // 在 user_login 组中使用 authMiddleware 中间件
	user.POST("/add", user_login.Adduser)
	user.DELETE("/del", user_login.DeleteUser)
	user.POST("/edit", user_login.UpdateUser)
}
