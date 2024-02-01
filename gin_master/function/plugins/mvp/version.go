package mvp

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Version_mvp(c *gin.Context) {
	var version Version_Mvp
	// 通过c.BindJSON将JSON请求正文绑定到version结构体
	if err := c.BindJSON(&version); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := "http://" + version.IP + "/mvp/version"
	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"error": response.StatusCode})
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "body": string(body)})
}
