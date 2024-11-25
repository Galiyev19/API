package handler

import (
	"API/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	adminAuth := router.Group("/admin/auth")
	{
		adminAuth.POST("sign-up", h.adminSignUp)
		adminAuth.POST("sign-in", h.adminSignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getUsers)
			users.GET("/:id", h.getUserByID)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return router
}
