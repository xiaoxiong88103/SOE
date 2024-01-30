package function

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Add_user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    string `json:"admin"`
}
type deluser struct {
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
