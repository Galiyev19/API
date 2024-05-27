package handlers

import (
	"fmt"
	"net/http"

	"API/internal/api/dto"
	"API/internal/api/helpers"
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

func (u *UserHandler) SignIn(c *gin.Context) {
	req := new(dto.RegisterUser)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse("Error bind json", false))
		return
	}

	userModel := model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err = u.service.User.Insert(userModel)
	if err != nil {
		fmt.Println("Error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse("Insert into DB", false))
		return
	}
	c.JSON(http.StatusCreated, helpers.GenerateResponse("test", true))
}
