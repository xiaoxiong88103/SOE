package mvp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"influxdb/config"
	"io"
	"io/ioutil"
	"net/http"
)

func Update_mvp(c *gin.Context) {
	var files Files
	if err := c.BindJSON(&files); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 校验参数是否为空
	if files.FileName == "" || files.IP == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数 Filename 和 IP 不能为空"})
		return
	}

	// 构建上传URL
	uploadURL := "http://" + files.IP + "/mvp/wget"
	prot, err := config.Dcode_json("web.json", "gin_prot")
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "打开文件出问题了", "error": err})
		return
	}

	data := map[string]string{
		"prot":     prot,
		"filename": files.FileName,
	}
	// 将 JSON 数据编码为字节数组
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON 编码错误:", err)
		return
	}
	req, err := http.NewRequest("POST", uploadURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("构建请求错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "出问题了", "error": err})
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusForbidden, gin.H{"msg": "出问题了", "error": "HTTP请求失败"})
		return
	}

	// 读取响应主体
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "出问题了", "error": err})
		return
	}
	// 打印响应内容
	fmt.Println("HTTP响应内容：", string(responseBody))

	// 文件上传成功后发送GET请求
	osURL := "http://" + files.IP + "/mvp/os?filename=/" + files.FileName
	responseOS, err := http.Get(osURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer responseOS.Body.Close()

	if responseOS.StatusCode != http.StatusOK {
		c.JSON(responseOS.StatusCode, gin.H{"error": fmt.Sprintf("HTTP status code: %d", responseOS.StatusCode)})
		return
	}

	responseOSBody, err := io.ReadAll(responseOS.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功", "response": string(responseBody), "message2": "更新mvp成功", "responseOS": string(responseOSBody)})

}
