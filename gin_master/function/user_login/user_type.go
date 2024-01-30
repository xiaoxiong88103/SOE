package user_login

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    string `json:"admin"`
	Group    string `json:"group"`
}

type Deluser struct {
	Username string `json:"username"`
}
