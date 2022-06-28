package handler

import (
	"books-api/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)
// @Summary Get all books
// @Security ApiKeyAuth
// @Tags books
// @Description get books
// @ID get-all-books
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /books/ [get]
func (h *Handler) getAll(c *gin.Context) {
	book, err := h.services.Book.GetAllBooks()
	if err != nil {
		logrus.Errorf("Error on getAll handler, %s", err)
		return
	}
	c.JSON(http.StatusOK, book)
}
// @Summary Get book By Id
// @Security ApiKeyAuth
// @Tags books
// @Description get book by id
// @ID get-book-by-id
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /books/:id [get]
func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	//TODO fix from Atoi to normal code
	newId, _ := strconv.Atoi(id)

	book, err := h.services.GetBookById(newId)
	if err != nil{
		logrus.Errorf("error during getting book by id, %s", err)
		return
	}

	c.JSON(http.StatusOK, book)
}
// @Summary Create book
// @Security ApiKeyAuth
// @Tags books
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body model.Book true "Book info"
// @Success 200 {integer} integer 1
// @Failure 400,404
// @Failure 500
// @Router /books/:id [post]
func (h *Handler) create(c *gin.Context) {
	book := model.Book{}

	err := c.BindJSON(&book)
	if err != nil {
		logrus.Errorf("error during binding JSON, %s", err)
		return
	}
	h.services.CreateBook(book)

	c.JSON(http.StatusOK, "Book was created successfully")
}
// @Summary Update book
// @Security ApiKeyAuth
// @Tags books
// @Description update book
// @ID update-book
// @Accept  json
// @Produce  json
// @Param input body model.Book true "Book info"
// @Success 200 {integer} integer 1
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /books/:id [patch]
func (h *Handler) update(c *gin.Context) {
	book := model.Book{}
	id := c.Param("id")

	err := c.BindJSON(&book)
	if err != nil{
		logrus.Errorf("error during binding JSON, %s", err)
		return
	}
	book.Id, _ = strconv.Atoi(id)

	h.services.UpdateBook(book)

	c.JSON(http.StatusOK, "Book was updated successfully")
}

// @Summary Delete book
// @Security ApiKeyAuth
// @Tags books
// @Description delete book
// @ID delete-book
// @Accept  json
// @Produce  json
// @Param input body model.Book true "Book info"
// @Success 200 {integer} integer 1
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /books/:id [delete]

func (h *Handler) delete(c *gin.Context) {
	//TODO delete govnocode
	id := c.Param("id")
	newId, _ := strconv.Atoi(id)

	err := h.services.DeleteBook(newId)
	if err != nil{
		logrus.Errorf("error during getting book by id, %s", err)
		return
	}
	c.JSON(http.StatusOK, "Book was deleted successfully")
}