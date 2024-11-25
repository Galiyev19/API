package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
}

func (h *Handler) getUsers(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": id,
	})
}

func (h *Handler) getUserByID(c *gin.Context) {
}

func (h *Handler) updateUser(c *gin.Context) {
}

func (h *Handler) deleteUser(c *gin.Context) {
}
