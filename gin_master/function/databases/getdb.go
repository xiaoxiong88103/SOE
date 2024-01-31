package databases

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"influxdb/config"
	"net/http"
	"time"
)

// @Summary 查询数据库
// @Description 查询数据库中的信息
// @Tags 数据库
// @Accept json
// @Produce json
// @Param Screen body Screen true "查询条件 timestart 写-10m or -10d -10s timestop写 now()代表当下时间"
// @Success 200 {object} Screen "查询成功" "返回的JSON数组" []QueryResult
// @Failure 400 {object} string "查询失败" "error":"xxx"
// @Router /show_db/screen [post]
func Get_db_screen(c *gin.Context) {
	var screenRequest Screen

	if err := c.ShouldBindJSON(&screenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := json_plus("url")
	token := json_plus("token")
	if url == "" || token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL or Token"})
		return
	}

	client := influxdb2.NewClient(url, token)
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create InfluxDB client"})
		return
	}
	defer client.Close()

	org := json_plus("org")
	master := json_plus("databases")
	name := screenRequest.Name
	timestart := screenRequest.Timestart
	timestop := screenRequest.Timestop

	query := fmt.Sprintf(`from(bucket: "%s")
    |> range(start: %s, stop: %s)
    |> filter(fn: (r) => r._measurement == "system_info" and r._field == "%s")`, master, timestart, timestop, name)

	fmt.Println("url:", url, "token", token)
	fmt.Println(query)

	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		fmt.Println("Query result is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query result is nil"})
		return
	}

	var queryResults []QueryResult
	for result.Next() {
		record := result.Record()
		if record == nil {
			fmt.Println("Record is nil")
			continue
		}

		// 假设 CPU 字段应该是 float64
		cpuValue, ok := record.Value().(float64)
		if !ok {
			fmt.Println("Failed to convert record value to float64")
			continue
		}

		localTime := record.Time().Local()
		queryResult := QueryResult{
			Time: localTime.String(),
			CPU:  cpuValue,
		}
		queryResults = append(queryResults, queryResult)
	}

	if queryErr := result.Err(); queryErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": queryErr.Error()})
		return
	}

	c.JSON(http.StatusOK, queryResults)
}

// Get_db_time 查询数据库时间范围内的数据
// @Summary 查询时间范围内的数据
// @Description 根据提供的时间范围查询 InfluxDB 中的 system_info 数据，并返回相关字段的值
// @Tags 数据库
// @Accept json
// @Produce json
// @Param screenRequest body Screen true "查询条件"
// @Success 200 {object} []Query_Result_all "查询成功，返回数据数组"
// @Failure 400 {object} string "查询参数错误" "返回的json({'error': 错误信息})"
// @Failure 500 {object} string "内部服务器错误" "返回的json({'error': 错误信息})"
// @Router /show_db/time [post]
func Get_db_time(c *gin.Context) {
	var screenRequest Screen

	if err := c.ShouldBindJSON(&screenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := json_plus("url")
	token := json_plus("token")
	if url == "" || token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL or Token"})
		return
	}

	client := influxdb2.NewClient(url, token)
	if client == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create InfluxDB client"})
		return
	}
	defer client.Close()

	org := json_plus("org")
	master := json_plus("databases")
	timestart := screenRequest.Timestart
	timestop := screenRequest.Timestop

	// 定义查询
	query1 := fmt.Sprintf(`
from(bucket: "%s")
  |> range(start: %s, stop: %s)
  |> filter(fn: (r) => r["_measurement"] == "system_info")
  |> filter(fn: (r) => r["_field"] == "cpu"  or r["_field"] == "gpu" or r["_field"] == "ioread" or r["_field"] == "iowrite" or r["_field"] == "netcon" or r["_field"] == "mem" or r["_field"] == "networkload" or r["_field"] == "networkup" or r["_field"] == "systemaver"  or r["_field"] == "vpu" or r["_field"] == "bandwidth")
  |> yield(name: "mean")`, master, timestart, timestop)

	// 执行第二次查询
	query2 := fmt.Sprintf(`from(bucket: "%s")
		|> range(start: %s, stop: %s)
		|> filter(fn: (r) => r._measurement == "system_info" and (r._field == "npu" or r._field == "disksize"))`, master, timestart, timestop)
	// 执行查询并处理结果
	var results []Query_Result_all
	if err := executeQuery(client, org, query1, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := executeQuery(client, org, query2, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回所有结果
	c.JSON(http.StatusOK, gin.H{"data": results, "code": "200"})

}

// executeQuery 执行查询并解析结果
func executeQuery(client influxdb2.Client, org string, query string, results *[]Query_Result_all) error {
	result, err := client.QueryAPI(org).Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("查询出错: %v", err)
	}
	defer result.Close()

	for result.Next() {
		record := result.Record()
		*results = append(*results, Query_Result_all{
			Time:  record.Time().Format(time.RFC3339),
			Field: record.Field(),
			Value: fmt.Sprintf("%v", record.Value()),
		})
	}

	if result.Err() != nil {
		return fmt.Errorf("处理查询结果时发生错误: %v", result.Err())
	}

	return nil
}

func json_plus(number string) string {
	par, err := config.Dcode_json("config.json", number)
	if err != nil {
		fmt.Println(err)
	}
	return par
}
