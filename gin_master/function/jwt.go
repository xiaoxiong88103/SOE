package function

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"influxdb/config"
	"net/http"
	"time"
)

var JwtKey []byte // 声明了一个全局变量用于存储 JWT 密钥

// Claims 结构体用于定义 JWT 的声明
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//全局变量
func init() {
	// 在应用程序初始化时，读取 JWT 密钥并存储在全局变量中
	tokenKey, err := config.Dcode_json("web.json", "token_key")
	if err != nil {
		fmt.Println("出错了:", err)
	}
	fmt.Println("读取key成功:", tokenKey)
	JwtKey = []byte(tokenKey)
}

// parseToken 用于解析 JWT token 并返回声明（claims）
func parseToken(tokenString string) (*Claims, error) {

	claims := &Claims{}

	// 解析 JWT token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("无效的 token")
	}

	return claims, nil
}

// generateValidToken 用于生成有效的 JWT token 以供测试使用
func generateValidToken(username string) string {

	time_token, err := config.DecodeJsonAsInt("web.json", "token_time_h")
	if err != nil {
		fmt.Println("出错了:", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time_token) * time.Hour).Unix(),
		},
	})

	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}

// authMiddleware 是一个中间件，用于验证 JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 token"})
			c.Abort()
			return
		}

		claims, err := parseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}

		// 在请求上下文中存储解析后的声明，以供后续处理函数使用
		c.Set("claims", claims)

		c.Next()
	}
}

// 解析 JWT 令牌并提取用户名
func DecodeToken_username(tokenString string) (string, error) {
	// 解析 JWT 令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法")
		}
		return JwtKey, nil // 使用相同的密钥进行解析
	})

	if err != nil {
		return "", err
	}

	// 检查令牌是否有效
	if !token.Valid {
		return "", fmt.Errorf("无效的令牌")
	}

	// 从令牌的声明中获取用户名
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("无法获取令牌声明")
	}

	username, ok := claims["Username"].(string)
	if !ok {
		return "", fmt.Errorf("无法获取用户名")
	}

	return username, nil
}
