package handler

import (
	"books-api/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	user := model.User{}

	err := c.BindJSON(&user)
	if err != nil{
		logrus.Errorf("Error on binding JSON, %s", err)
		return
	}
	id, err  := h.services.Create(user)
	if err != nil {
		logrus.Errorf("Error on creating user, %s", err)
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	user := model.User{}

	err := c.BindJSON(&user)
	if err != nil {
		logrus.Errorf("Error on binding JSON, %s", err)
		return
	}
	token, err := h.services.Get(user.Name, user.Password)
	if err != nil {
		logrus.Errorf("Error on getting user data, %s", err)
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": token.Id,
	})
}
