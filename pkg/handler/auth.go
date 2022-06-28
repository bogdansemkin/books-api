package handler

import (
	"books-api/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	user := model.User{}

	err := c.BindJSON(&user)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err  := h.services.Authorization.Create(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	token, err := h.services.Authorization.Get(user.Name, user.Password)
	if err != nil {
		logrus.Errorf("Error on getting user data, %s", err)
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": token.Id,
	})
}
