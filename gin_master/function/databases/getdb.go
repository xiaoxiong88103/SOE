package databases

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"influxdb/config"
	"net/http"
)

func Get_db_screen(c *gin.Context) {
	var screenRequest Screen

	if err := c.ShouldBindJSON(&screenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := influxdb2.NewClient(json_plus("url"), json_plus("token"))
	org := json_plus("org")
	master := json_plus("databases")
	name := screenRequest.Name // 获取名为 "name" 的值
	time := screenRequest.Time // 获取名为 "time" 的值

	// 创建查询
	query := fmt.Sprintf(`from(bucket: "%s")
    |> range(start: "%s")
    |> filter(fn: (r) => r._measurement == "system_info" and r._field == "%s")`, master, time, name)

	// 执行查询
	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 处理查询结果
	var queryResults []QueryResult // 用于存储查询结果的切片

	for result.Next() {
		if result.Record().Field() == name {
			// 转换时间为本地时区
			localTime := result.Record().Time().Local()

			// 创建 QueryResult 结构体并添加到切片中
			queryResult := QueryResult{
				Time: localTime.String(),
				CPU:  result.Record().Value(),
			}

			queryResults = append(queryResults, queryResult)
		}
	}

	// 检查查询过程中是否有错误发生
	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Err().Error()})
		return
	}

	// 返回查询结果切片作为 JSON 响应
	c.JSON(http.StatusOK, queryResults)
}

func json_plus(number string) string {
	par, err := config.Dcode_json("config.json", number)
	if err != nil {
		fmt.Println(err)
	}
	return par
}
