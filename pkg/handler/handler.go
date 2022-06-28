package handler

import (
	"books-api/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "books-api/docs"
)

type Handler struct {
	services *service.Service
}
func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	books := router.Group("/books")
	{
		books.GET("", h.getAll)
		books.GET("/:id", h.getById)
		books.POST("", h.create)
		books.PATCH("/:id", h.update)
		books.DELETE("/:id", h.delete)
	}

	return router
}
