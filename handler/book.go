package handler

import (
	"books-api/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getAll(c *gin.Context) {
	book, err := h.services.GetAllBooks()
	if err != nil {
		log.Fatalf("Error on getAll handler, %s", err)
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) getById(c *gin.Context) {

}

func (h *Handler) create(c *gin.Context) {
	book := model.Book{}

	err := c.BindJSON(&book)
	if err != nil {
		log.Fatalf("error during binding JSON, %s", err)
		return
	}
	h.services.CreateBook(book)

}

func (h *Handler) update(c *gin.Context) {

}

func (h *Handler) delete(c *gin.Context) {

}