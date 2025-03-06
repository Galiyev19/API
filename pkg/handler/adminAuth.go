package handler

import (
	"net/http"

	"API/pkg/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
// @Tags		Admin
// @Accept 		json
// @Produce 	json
// @Param		input	body	models.AdminRequest	true	"Admin credentials"
// @Success		200		{object}	map[string]string	"token"
// @Failure		400		{object}	map[string]string	"Bad Request"
// @Failure		500		{object}	map[string]string	"Internal Server Error"
// @Router			/admin/auth/sign-in [post]
func (h *Handler) AdminSignIn(c *gin.Context) {
	var input models.AdminRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	admin, err := h.service.GetAdmin(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if admin == nil {
		newErrorResponse(c, http.StatusNotFound, "Admin not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid password")
		return
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
// func (h *Handler) TestRoute(c *gin.Context) {
// 	c.JSON(http.StatusOK, "OK")
// }
