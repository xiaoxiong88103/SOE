package Network

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Nmap_node 处理端口扫描请求
// @Summary 端口扫描操作
// @Description 接收一个IP地址和一个端口列表，使用nmap执行端口扫描并返回扫描结果。
// @Tags 基础功能
// @Accept json
// @Produce json
// @Param ip query string true "IP地址"
// @Param ports query string true "端口列表，端口之间用逗号分隔，如80,443,8080 or all"
// @Success 200 {array} string "成功响应，返回扫描结果数组"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部服务器错误"
// @Router /basic/nmap [get]
func Nmap_node(c *gin.Context) {
	ip := c.Query("ip")
	portsQuery := c.Query("ports") // 期望格式为 "80,443,8080 or all"

	if ip == "" || portsQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IP or ports parameter missing"})
		return
	}

	var ports []string
	if portsQuery == "all" {
		// 如果portsQuery为"all"，则设置ports为"1-65535"，或者根据需要生成完整的端口列表
		ports = []string{"1-65535"}
	} else {
		// 否则，按照逗号分割portsQuery
		ports = strings.Split(portsQuery, ",")
	}

	// 设置 Nmap 扫描选项
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(ports...),
		nmap.WithContext(ctx),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create nmap scanner", "detail": err.Error()})
		return
	}

	// 运行扫描
	result, warnings, err := scanner.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to run nmap scan", "detail": err.Error()})
		return
	}
	if warnings != nil {
		fmt.Println("Warnings: ", warnings)
	}

	// 解析并返回扫描结果
	var scanResults []map[string]interface{}
	for _, host := range result.Hosts {
		for _, port := range host.Ports {
			scanResult := map[string]interface{}{
				"ip":       host.Addresses[0],
				"port":     port.ID,
				"state":    port.State.State,
				"service":  port.Service.Name,
				"protocol": port.Protocol,
			}
			scanResults = append(scanResults, scanResult)
		}
	}

	c.JSON(http.StatusOK, gin.H{"results": scanResults})
}
