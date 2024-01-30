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
