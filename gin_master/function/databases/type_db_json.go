package databases

//单独查询json
type Screen struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

//单独查询返回的json
type QueryResult struct {
	Time string      `json:"time"`
	CPU  interface{} `json:"cpu"`
}
