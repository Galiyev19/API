package handlers

import (
	"net/http"

	"API/internal/api/helpers"
	"API/internal/config"

	"github.com/gin-gonic/gin"
)

type (
	HealthHandler struct {
		cfg *config.Config
	}
	envelope map[string]interface{}
)

func NewHealthHandler(cfg *config.Config) *HealthHandler {
	return &HealthHandler{
		cfg: cfg,
	}
}

func (h *HealthHandler) Health(c *gin.Context) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"enviroment": h.cfg.Enviroment,
			"version":    h.cfg.Version,
		},
	}
	c.JSON(http.StatusOK, helpers.GenerateResponse(env, true))
}
