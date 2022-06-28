package handler

import (
	"books-api/model"
	"books-api/service"
	mock_service "books-api/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_sigUp(t *testing.T){
	type mockBehavior func(s *mock_service.MockAuthorization, user model.User)

	testTable := []struct{
		name 				string
		inputBody 			string
		inputUser 			model.User
		mockBehavior 		mockBehavior
		expectedStatusCode 	int
		expectedRequestBody string
	} {
		{
			name : "OK",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: model.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User){
				s.EXPECT().Create(user).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: `{"id":1}`,
		},
		/*{
			name : "Empty fields",
			inputBody: `{"username":"test"}`,
			inputUser: model.User{},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User){},
			expectedStatusCode: 400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},*/
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: model.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
				r.EXPECT().Create(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedRequestBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable{
		t.Run(testCase.name, func(t *testing.T){
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization:auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			//test Request

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			//perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}