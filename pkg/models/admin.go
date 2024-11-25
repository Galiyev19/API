package models

import "time"

type Admin struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role" binding:"required"`
}

type AdminRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
