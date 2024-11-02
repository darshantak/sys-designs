package rest

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChatRoom struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

