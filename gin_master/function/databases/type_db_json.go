package databases

//单独查询json
type Screen struct {
	Name      string `json:"name"`
	Timestart string `json:"timestart"`
	Timestop  string `json:"timestop"`
}

//单独查询返回的json
type QueryResult struct {
	Time string      `json:"time"`
	CPU  interface{} `json:"system"`
}

type Query_Result_all struct {
	Time  string
	Field string
	Value string
}
