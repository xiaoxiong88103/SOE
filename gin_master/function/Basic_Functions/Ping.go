package Basic_Functions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
	"net/http"
	"time"
)

// Ping_node 处理ping请求
// @Summary Ping操作
// @Description 发送ping到指定的IP地址，返回延迟情况
// @Tags 网络
// @Accept  json
// @Produce  json
// @Param   ip   query    string     true  "IP地址"
// @Success 200 {object} PingResult "成功响应"
// @Failure 400 {object} string "请求参数错误"
// @Failure 500 {object} string "内部服务器错误"
// @Router /basic/ping [get]
func Ping_node(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IP地址参数缺失"})
		return
	}

	pinger, err := ping.NewPinger(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("创建ping请求失败: %v", err)})
		return
	}
	pinger.Count = 3                  // 发送ping的次数
	pinger.Timeout = time.Second * 10 // 超时时间
	pinger.SetPrivileged(true)        // 在unix系统上需要root权限

	// 执行ping操作
	err = pinger.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("执行ping失败: %v", err)})
		return
	}

	stats := pinger.Statistics() // 获取统计结果
	c.JSON(http.StatusOK, gin.H{
		"ip":           ip,                    // 目标 IP 地址，是 `ping` 操作的对象
		"rtt":          stats.AvgRtt.String(), // "rtt" 代表往返时间（Round-Trip Time），这是 `ping` 命令发送数据到接收回应所花费的平均时间。这里的值是平均往返时间的字符串表示。
		"packets_sent": stats.PacketsSent,     // "packets_sent" 表示发送的数据包数量。这是执行 `ping` 操作时，实际尝试发送到目标 IP 地址的 ICMP 数据包总数。
		"packets_recv": stats.PacketsRecv,     // "packets_recv" 表示接收到的数据包数量。这是从目标 IP 地址成功接收到的 ICMP 响应数据包总数。如果这个数值小于 "packets_sent"，可能表示网络中存在丢包。
	})

}
