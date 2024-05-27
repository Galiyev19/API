package handlers

import (
	"net/http"

	"API/internal/api/helpers"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) SignIn(c *gin.Context) {
	c.JSON(http.StatusCreated, helpers.GenerateResponse("test", true))
}
