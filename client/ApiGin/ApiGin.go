package ApiGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"influxdb/client/ApiGin/function"
	"influxdb/config"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApiGin() {
	route := gin.Default()
	gin.DefaultWriter = ioutil.Discard
	//获取白名单IP和 接口用到的端口的函数
	ip, prot := getconfig()

	// 添加 IP 白名单中间件
	route.Use(IPWhiteList([]string{ip}))

	route.GET("/ping", function.Get_Null)

	route.Run(":" + prot)
}

func IPWhiteList(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的 IP 地址
		ip := c.ClientIP()
		// 检查 IP 地址是否在白名单中
		allowed := false
		for _, value := range whitelist {
			if value == ip {
				allowed = true
				break
			}
		}
		// 如果 IP 地址不在白名单中，则返回错误信息
		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "IP address not allowed"})
			return
		}
		// 允许请求继续访问后续的处理函数
		c.Next()
	}
}

// 获取配置里的函数
func getconfig() (string, string) {
	//获取验证的IP
	ip, err := config.Dcode_json("config.json", "serverip")
	if err != nil {
		fmt.Println("获取IP错误:", err)
	}
	parts := strings.Split(ip, ":")

	//获取端口号
	prot, err := config.Dcode_json("config.json", "gin_prot")
	if err != nil {
		fmt.Println("开启的时候报错:", err)
	}
	return parts[0], prot
}
