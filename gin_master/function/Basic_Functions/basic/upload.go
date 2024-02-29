package basic

import (
	"github.com/gin-gonic/gin"
	"influxdb/config"
	"net/http"
	"os"
	"path/filepath"
)

var pathfiles = config.Path_files

// UploadFile 上传文件接口
// @Summary 上传文件
// @Description 用于上传文件的接口
// @Tags 文件操作
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "要上传的文件"
// @Success 200 {object} string "{"message": "文件上传成功"}"
// @Failure 400 {object} string "{"error": "错误信息"}"
// @Failure 500 {object} string "{"error": "内部服务器错误"}"
// @Router /files/upload [put]
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将上传的文件保存到指定目录
	err = c.SaveUploadedFile(file, pathfiles+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功"})
}

// RemoveFile 删除文件接口
// @Summary 删除文件
// @Description 用于删除文件的接口
// @Tags 文件操作
// @Accept json
// @Produce json
// @Param filename path string true "要删除的文件名"
// @Success 200 {object} string "{"message": "文件删除成功"}"
// @Failure 404 {object} string "{"error": "文件不存在"}"
// @Failure 500 {object} string "{"error": "内部服务器错误"}"
// @Router /files/{filename} [delete]
func RemoveFile(c *gin.Context) {
	filename := c.Param("filename")

	// 构建文件路径
	filePath := pathfiles + filename

	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "文件不存在"})
		return
	}

	// 删除文件
	err = os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件删除成功"})
}

// GetFilelist 获取文件目录列表
// @Summary 获取文件目录列表
// @Description 提供访问'/files'目录下所有文件的列表。
// @Tags 文件操作
// @Accept json
// @Produce json
// @Success 200 {array} string "返回文件目录下的文件列表"
// @Router /files/list [get]
func GetFilelist(c *gin.Context) {
	var files []string

	// 遍历目录，列出所有文件
	err := filepath.Walk(pathfiles, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 如果你只想要文件名，而不是完整路径，可以使用 filepath.Base(path)
			files = append(files, filepath.Base(path))
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}
