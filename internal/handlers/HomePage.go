package handlers

import (
	"github.com/gin-gonic/gin"
)

type envelope map[string]interface{}

func (h *Handlers) HomePage(c *gin.Context) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"enviroment": h.cfg.Enviroment,
			"version":    h.cfg.Version,
		},
	}
	c.IndentedJSON(200, env)
}
