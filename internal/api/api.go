package api

import (
	"log"

	"API/internal/api/router"
	"API/internal/config"
	"API/internal/service"

	"github.com/gin-gonic/gin"
)

type Api struct {
	service *service.Service
	cfg     *config.Config
}

func NewApi(service *service.Service, cfg *config.Config) *Api {
	return &Api{
		service: service,
		cfg:     cfg,
	}
}

func (h *Api) InitServer() {
	r := gin.New()

	h.InitRoutes(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("run server error = %v", err)
	}
}

func (h *Api) InitRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Health Check
		health := v1.Group("/health")

		// Health Check
		router.Health(health, h.cfg)

		// User
		user := v1.Group("/users")

		router.User(user, h.service)

	}
}
