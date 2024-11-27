package handler

import (
	_ "API/docs"
	"API/pkg/service"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

// @title API Documentation
// @version 1.0
// @description This is a simple API documentation example
// @host localhost:8080
// @BasePath /api
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

		products := api.Group("/products")
		{
			products.POST("/", h.CreateProduct)
			products.GET("/", h.GetProducts)
			products.GET("/:id")
			products.PUT("/:id")
			products.DELETE("/:id")
		}
	}

	return router
}
