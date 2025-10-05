package models

type User struct {
	ID       int64
	Username string
	Password string
}

type UserData struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
