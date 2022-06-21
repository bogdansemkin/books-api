package handler

import (
	"books-api/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	user := model.User{}

	c.BindJSON(&user)
	h.services.Create(user)
}

func (h *Handler) signIn(c *gin.Context) {

}
