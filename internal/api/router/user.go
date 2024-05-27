package router

import (
	"API/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	h := handlers.NewUserHandler()

	r.POST("/sign-in", h.SignIn)
}
