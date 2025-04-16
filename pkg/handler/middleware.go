package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token := strings.TrimSpace(header)

	id, role, err := h.service.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("id", id)
	c.Set("role", role)
}

func (h *Handler) adminIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token := strings.TrimSpace(header)
	id, role, err := h.service.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "Access denied")
		return
	}

	c.Set("id", id)
	c.Set("role", role)
}
