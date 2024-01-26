package function

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login_web(c *gin.Context) {

	data := gin.H{"code": http.StatusOK}
	c.JSON(http.StatusOK, data)

}
