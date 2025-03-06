package handler

import (
	"API/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     Create User
// @Security    BearerAuth
// @Description Creates a new user
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       input body models.UserRequest true "User credentials"
// @Success     200 {object} map[string]string "message"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/users/create-user [post]
func (h *Handler) createUser(c *gin.Context) {
	var input models.UserRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Email == "" || input.Password == "" || input.UserName == "" {
		newErrorResponse(c, http.StatusBadRequest, "Email, UserName and Password are required")
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
		UserName: input.UserName,
	}
	_, err := h.service.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Status Internal Server")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created",
	})

}

// @Summary     Get User List
// @Security    BearerAuth
// @Description Get users list
// @Tags        Users
// @Accept      json
// @Produce     json
// @Success 200 {object} map[string]string "message"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/users/user-list [get]
func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.service.Users.GetListUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Status Internal Server")
		return
	}
	fmt.Println("users", users)
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

func (h *Handler) getUserByID(c *gin.Context) {
}

func (h *Handler) updateUser(c *gin.Context) {
}

func (h *Handler) deleteUser(c *gin.Context) {
}
