package handlers

import (
	"net/http"

	"API/internal/api/dto"
	"API/internal/api/helpers"
	"API/internal/api/validator"
	"API/internal/model"
	"API/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.Service
}

func NewUserHandler(s *service.Service) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// REGISTER USER ROUTER
func (u *UserHandler) SignUp(c *gin.Context) {
	req := new(dto.RegisterUser)
	err := c.ShouldBindJSON(&req)

	// check valid email and password
	v := validator.NewValidator()
	dto.ValidateUser(v, req)

	if err != nil || !v.Valid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse(v.Errors, false))
		return
	}

	userModel := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err = u.service.User.Insert(userModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse(err.Error(), false))
		return
	}
	c.JSON(http.StatusCreated, helpers.GenerateResponse("test", true))
}
