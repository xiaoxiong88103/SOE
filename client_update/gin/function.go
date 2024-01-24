package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get_Null(c *gin.Context) {
	data := gin.H{"code": http.StatusOK}
	c.JSON(http.StatusOK, data)
}
