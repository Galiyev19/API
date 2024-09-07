package router

import (
	"API/internal/api/handlers"
	"API/internal/config"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewHealthHandler(cfg)

	r.GET("/", handler.Health)
}
