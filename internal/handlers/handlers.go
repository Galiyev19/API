package handlers

import (
	"API/internal/config"
	"API/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service *service.Service
	cfg     *config.Config
}

func NewHandler(service *service.Service, cfg *config.Config) *Handlers {
	return &Handlers{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", h.HomePage)

	return router
}
