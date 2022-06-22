package handler

import (
	"books-api/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getAll(c *gin.Context) {
	book, err := h.services.GetAllBooks()
	if err != nil {
		log.Fatalf("Error on getAll handler, %s", err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	//TODO fix from Atoi to normal code
	newId, _ := strconv.Atoi(id)

	book, err := h.services.GetBookById(newId)
	if err != nil{
		log.Fatalf("error during getting book by id, %s", err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) create(c *gin.Context) {
	book := model.Book{}

	err := c.BindJSON(&book)
	if err != nil {
		log.Fatalf("error during binding JSON, %s", err)
		return
	}
	h.services.CreateBook(book)

	c.JSON(http.StatusOK, "User was created successfully")
}

func (h *Handler) update(c *gin.Context) {

}

func (h *Handler) delete(c *gin.Context) {
	//TODO delete govnocode
	id := c.Param("id")
	newId, _ := strconv.Atoi(id)

	err := h.services.DeleteBook(newId)
	if err != nil{
		log.Fatalf("error during getting book by id, %s", err)
		return
	}
	c.JSON(http.StatusOK, "Book was deleted successfully")

}