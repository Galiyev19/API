package handler

import (
	"net/http"

	"API/pkg/models"

	"github.com/gin-gonic/gin"
)

// @Summary		admin sign-up
// @Description	Just a test route to check Swagger generation
// @Tags		Admin
// @Accept 		json
// @Produce 	json
// @Param		input	body	models.AdminRequest	true	"Admin credentials"
// @Success		200		{object}	map[string]string	"token"
// @Failure		400		{object}	map[string]string	"Bad Request"
// @Failure		500		{object}	map[string]string	"Internal Server Error"
// @Router		/admin/auth/sign-up [post]
func (h *Handler) AdminSignUp(c *gin.Context) {
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

// @Summary		admin sign-in
// @Description	Just a test route to check Swagger generation
// @Tags			Test
// @Success		200	{string}	string	"OK"
// @Router			/admin/auth/sign-in [post]
func (h *Handler) AdminSignIn(c *gin.Context) {
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

// @Summary		Test route
// @Description	Just a test route to check Swagger generation
// @Tags			Test
// @Success		200	{string}	string	"OK"
// @Router			/admin/auth/test [post]
func (h *Handler) TestRoute(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
