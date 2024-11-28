package models

type SuccessResponse struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}
