package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

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
// @Router /files/upload [post]
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将上传的文件保存到指定目录
	err = c.SaveUploadedFile(file, "./files/"+file.Filename)
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
	filePath := "./files/" + filename

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
