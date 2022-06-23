package handler

import (
	"books-api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}
func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

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
