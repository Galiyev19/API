package handler

import (
	"net/http"

	"API/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) adminSignUp(c *gin.Context) {
	var input models.Admin

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.Authorization.CreateAdmin(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) adminSignIn(c *gin.Context) {
	var input models.AdminRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.service.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
