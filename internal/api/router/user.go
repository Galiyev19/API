package router

import (
	"API/internal/api/handlers"
	"API/internal/service"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup, service *service.Service) {
	h := handlers.NewUserHandler(service)

	r.POST("/sign-up", h.SignUp)
}
