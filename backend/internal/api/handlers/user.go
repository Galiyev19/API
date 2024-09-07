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

	if err != nil && !v.Valid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse(v.Errors, false))
		return
	}

	// check user email is exist in DB
	existUser, _ := u.service.User.GetUserByEmail(req.Email)

	if existUser.Email == req.Email {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse("This email is exist", false))
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
	c.JSON(http.StatusCreated, helpers.GenerateResponse("SUCESS - user created", true))
}

// LOGIN USER

func (u *UserHandler) SignIn(c *gin.Context) {
	req := new(dto.RegisterUser)
	err := c.ShouldBindJSON(&req)

	// check valid email and password
	v := validator.NewValidator()
	dto.ValidateUser(v, req)

	// error json data
	if err != nil && !v.Valid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateResponse(v.Errors, false))
		return
	}

	// generate token
	token, err := u.service.User.GenerateToken(req.Email, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// return generated token
	c.JSON(http.StatusOK, helpers.GenerateResponse(map[string]interface{}{
		"token": token,
	}, true))
}
