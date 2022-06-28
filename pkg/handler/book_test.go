package handler
//
//import (
//	"books-api/model"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/magiconair/properties/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func SetUpRouter() *gin.Engine{
//	router := gin.Default()
//	return router
//}
//
//func TestGetAllBooks(t *testing.T){
//	r := SetUpRouter()
//	handler := Handler{}
//	r.GET("/books", handler.getAll)
//	req, _ := http.NewRequest("GET", "/books", nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//
//	var books []model.Book
//	json.Unmarshal(w.Body.Bytes(), &books)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//}