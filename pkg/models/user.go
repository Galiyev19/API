package models

type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"created_at"`
}
