package user_ini

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

// UserSections 包含不同部分的用户信息
type UserSections struct {
	User  map[string]string `json:"user"`
	Admin map[string]string `json:"admin"`
	Group map[string]string `json:"group"`
}
