package Plugins

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func WgetMVP(c *gin.Context) {
	var jsonInput WgetJson
	clientIP := c.ClientIP()
	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 下载文件
	response, err := http.Get("http://" + clientIP + ":" + jsonInput.Prot + "/files/" + jsonInput.Filename)
	fmt.Println("http://" + clientIP + ":" + jsonInput.Prot + "/files/" + jsonInput.Filename)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"message": fmt.Sprintf("HTTP status code: %d", response.StatusCode)})
		return
	}

	// 创建目标文件
	filename := filepath.Join("/", jsonInput.Filename)
	file, err := os.Create(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 将文件内容复制到本地文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "200"})
}

func MVPOS(c *gin.Context) {
	filename := c.Query("filename") // 从查询参数中获取要下载的文件的URL
	//先停止mvp
	startStopCommand := fmt.Sprintf("/opt/MVP%s/start stop", GetSysver())
	cmd := exec.Command("sh", "-c", startStopCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tarCommand := fmt.Sprintf("tar -zxvf %s -C /", filename)

	// 执行解压命令
	tarCmd := exec.Command("sh", "-c", tarCommand)
	tarCmd.Stdout = os.Stdout
	tarCmd.Stderr = os.Stderr
	if err := tarCmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 执行4次 mysql 命令
	for i := 0; i < 4; i++ {
		mysqlCommand := fmt.Sprintf("mysql -D VAS < /opt/MVP%s/update.sql -f", GetSysver())
		mysqlCmd := exec.Command("sh", "-c", mysqlCommand)
		mysqlCmd.Stdout = os.Stdout
		mysqlCmd.Stderr = os.Stderr
		if err := mysqlCmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// 执行 ldconfig 命令两次，每次间隔1秒
	for i := 0; i < 2; i++ {
		ldconfigCmd := exec.Command("ldconfig")
		ldconfigCmd.Stdout = os.Stdout
		ldconfigCmd.Stderr = os.Stderr
		if err := ldconfigCmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		time.Sleep(1 * time.Second)
	}

	// 开启 MVP，执行2次，每次间隔1秒
	for i := 0; i < 2; i++ {
		startCommand := fmt.Sprintf("/opt/MVP%s/start", GetSysver())
		startCmd := exec.Command("sh", "-c", startCommand)
		startCmd.Stdout = os.Stdout
		startCmd.Stderr = os.Stderr
		if err := startCmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		time.Sleep(3 * time.Second)
	}

	remove := exec.Command("rm", "-rf", "/"+filename)
	remove.Stdout = os.Stdout
	remove.Stderr = os.Stderr
	if err := remove.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200"})

}

func GetSysver() string {
	arch := runtime.GOARCH
	if arch == "amd64" {
		return "64"
	} else if arch == "arm64" {
		return "arm"
	} else {
		return ""
	}
}

func GetVersionTxt(c *gin.Context) {
	// 读取文件内容
	content, err := ioutil.ReadFile("/opt/MVParm/version.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将文件内容返回给前端
	c.String(http.StatusOK, string(content))
}
