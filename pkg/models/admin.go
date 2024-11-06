package models

type Admin struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	Role      string `json:"role"`
}
