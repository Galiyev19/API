package handler

import (
	"net/http"

	"API/pkg/models"

	"github.com/gin-gonic/gin"
)

// SignUp Admin
// @Summary Sign-Up Admin
// @Description Sign-up Admin and get JWT TOKEN
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body models.AdminRequest true "Admin information"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.Error
// @Router /admin/auth/sign-up [post]

func (h *Handler) adminSignUp(c *gin.Context) {
	var input models.AdminRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Email == "" || input.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "Email and Password are required")
		return
	}

	admin := models.Admin{
		Email:    input.Email,
		Password: input.Password,
		Role:     "admin",
	}

	_, err := h.service.Authorization.CreateAdmin(admin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// SignUp Admin
// @Summary Sign-In Admin
// @Description Sign-In Admin and get JWT TOKEN
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body models.AdminRequest true "Admin information"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.Error
// @Router /admin/auth/sign-in [post]

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
