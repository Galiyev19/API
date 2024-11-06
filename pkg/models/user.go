package models

type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
