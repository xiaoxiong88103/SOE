package function

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginWeb 处理用户登录请求
// @Summary 用户登录
// @Description 处理用户登录请求，验证用户名和密码。
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param user body User true "登录信息"
// @Success 200 {object} string "成功登录，返回token" 返回的json({"message": "用户名和密码验证成功", "token":token })
// @Failure 401 {object} string "认证失败" 返回的json({"error": "验证失败", "User": 用户名})
// @Router /login [post]
func LoginWeb(c *gin.Context) {
	var userJSON User

	// 尝试解析请求中的 JSON 数据到 userJSON 结构体
	if err := c.ShouldBindJSON(&userJSON); err == nil {
		// 验证用户名和密码
		if validateUser(userJSON.Username, userJSON.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "用户名和密码验证成功", "token": generateValidToken(userJSON.Username)})
			return
		}
	}

	// 验证失败的情况下返回错误响应
	c.JSON(http.StatusUnauthorized, gin.H{"error": "验证失败", "User": userJSON})
}

// 添加用户的接口
// @Summary 用户新增
// @Description 用户进来后增加其他人的权限 账号密码等开通
// @Tags 用户新增
// @Accept json
// @Produce json
// @Param Add_user body Add_user true "新增"
// @Success 200 {object} string "添加成功" 返回的json({"message": "用户添加成功"})
// @Failure 401 {object} string "添加失败" 返回的json({"error": error , "massage": "报错消息"})
// @Router /user/add [post]
func Adduser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// 解析 JWT 令牌并提取用户名
	username, err := DecodeToken_username(tokenString)
	if err != nil {
		// 处理解析错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
		return
	}
	user_admin := Show_admin(username)
	if user_admin <= "90" {
		c.JSON(http.StatusForbidden, gin.H{"massage": "对不起你权限不足"})
		return
	} else {
		// 解析请求中的 JSON 数据到 adduser 结构体
		var userData Add_user
		if err := c.BindJSON(&userData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 调用添加用户数据的函数
		err := addUserdata(userData.Username, userData.Password, userData.Admin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "用户添加成功"})
	}

}
