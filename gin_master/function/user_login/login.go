package user_login

import (
	"github.com/gin-gonic/gin"
	"influxdb/gin_master/function"
	"influxdb/gin_master/function/user_ini"
	"net/http"
)

// LoginWeb 处理用户登录请求
// @Summary 用户登录
// @Description 处理用户登录请求，验证用户名和密码。
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param user_login body User true "登录信息"
// @Success 200 {object} string "成功登录，返回token" 返回的json({"message": "用户名和密码验证成功", "token":token })
// @Failure 401 {object} string "认证失败" 返回的json({"error": "验证失败", "User": 用户名})
// @Router /login [post]
func LoginWeb(c *gin.Context) {
	var userJSON Login

	// 尝试解析请求中的 JSON 数据到 userJSON 结构体
	if err := c.ShouldBindJSON(&userJSON); err == nil {
		// 验证用户名和密码
		if user_ini.ValidateUser(userJSON.Username, userJSON.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "用户名和密码验证成功", "token": function.GenerateValidToken(userJSON.Username)})
			return
		}
	}

	// 验证失败的情况下返回错误响应
	c.JSON(http.StatusUnauthorized, gin.H{"error": "验证失败", "User": userJSON})
}

// 添加用户的接口
// @Summary 用户新增
// @Description 用户进来后增加其他人的权限 账号密码等开通
// @Tags 用户
// @Accept json
// @Produce json
// @Param User body User true "新增"
// @Success 200 {object} string "添加成功" 返回的json({"message": "用户添加成功"})
// @Failure 401 {object} string "添加失败" 返回的json({"error": error , "massage": "报错消息"})
// @Router /user_login/add [post]
func Adduser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// 解析 JWT 令牌并提取用户名
	username, err := function.DecodeToken_username(tokenString)
	if err != nil {
		// 处理解析错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
		return
	}
	user_admin := user_ini.Show_admin(username)
	if user_admin <= "90" {
		c.JSON(http.StatusForbidden, gin.H{"massage": "对不起你权限不足"})
		return
	} else {
		// 解析请求中的 JSON 数据到 adduser 结构体
		var userData User
		if err := c.BindJSON(&userData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 调用添加用户数据的函数
		err := user_ini.AddUserdata(userData.Username, userData.Password, userData.Admin, userData.Group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "用户添加成功"})
	}

}

// 删除用户的接口
// @Summary 用户删除
// @Description 删除指定的用户账号及相关权限
// @Tags 用户
// @Accept json
// @Produce json
// @Param username path string true "用户名"
// @Success 200 {object} string "删除成功" 返回的json({"message": "用户删除成功"})
// @Failure 401 {object} string "删除失败" 返回的json({"error": error , "message": "报错消息"})
// @Router /user_login/delete/{username} [delete]
func DeleteUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// 解析 JWT 令牌并提取用户名
	username, err := function.DecodeToken_username(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
		return
	}

	// 检查当前用户是否有足够权限执行删除操作
	user_admin := user_ini.Show_admin(username)
	if user_admin <= "90" {
		c.JSON(http.StatusForbidden, gin.H{"message": "对不起你权限不足"})
		return
	}

	// 从路径参数中获取要删除的用户名
	usernameToDelete := c.Param("username")

	// 调用删除用户数据的函数
	err = user_ini.RemoveUserdata(usernameToDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

// 修改用户的接口
// @Summary 用户修改
// @Description 修改指定的用户账号及相关权限
// @Tags 用户
// @Accept json
// @Produce json
// @Param username path string true "用户名"
// @Success 200 {object} User "删除成功"
// @Failure 401 {object} string "删除失败" 返回的json({"error": error , "message": "报错消息"})
// @Router /user_login/edti [POST]
func UpdateUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// 解析 JWT 令牌并提取用户名
	username, err := function.DecodeToken_username(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
		return
	}

	// 检查当前用户是否有足够权限执行删除操作
	user_admin := user_ini.Show_admin(username)
	if user_admin <= "90" {
		c.JSON(http.StatusForbidden, gin.H{"message": "对不起你权限不足"})
		return
	}

	var updateuser User
	if err := c.BindJSON(&updateuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用删除用户数据的函数
	err = user_ini.UpdateUserdata(updateuser.Username, updateuser.Password, updateuser.Admin, updateuser.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户修改成功"})
}
