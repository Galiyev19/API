package model

type User struct {
	UserId    string
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string
}
