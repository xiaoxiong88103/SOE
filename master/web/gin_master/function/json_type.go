package function

type login struct {
	username string `json:"username"`
	password string `json:"password"`
}

type adduser struct {
	login
	admin int64 `json:"admin"`
}
type deluser struct {
	login
}

type addipgroup struct {
	ipv4  string `json:"ipv_4"`
	ipv6  string `json:"ipv_6"`
	group string `json:"group"`
	name  string `json:"name"`
	notes string `json:"notes"`
}

type editipgrop struct {
	addipgroup
}

type delipgrop struct {
	addipgroup
}
