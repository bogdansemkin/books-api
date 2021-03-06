// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "books-api/pkg/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthorization) Create(user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthorizationMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorization)(nil).Create), user)
}

// Get mocks base method.
func (m *MockAuthorization) Get(name, password string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name, password)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAuthorizationMockRecorder) Get(name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAuthorization)(nil).Get), name, password)
}

// MockBook is a mock of Book interface.
type MockBook struct {
	ctrl     *gomock.Controller
	recorder *MockBookMockRecorder
}

// MockBookMockRecorder is the mock recorder for MockBook.
type MockBookMockRecorder struct {
	mock *MockBook
}

// NewMockBook creates a new mock instance.
func NewMockBook(ctrl *gomock.Controller) *MockBook {
	mock := &MockBook{ctrl: ctrl}
	mock.recorder = &MockBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBook) EXPECT() *MockBookMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBook) CreateBook(book model.Book) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", book)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookMockRecorder) CreateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBook)(nil).CreateBook), book)
}

// DeleteBook mocks base method.
func (m *MockBook) DeleteBook(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookMockRecorder) DeleteBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBook)(nil).DeleteBook), id)
}

// GetAllBooks mocks base method.
func (m *MockBook) GetAllBooks() ([]model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks")
	ret0, _ := ret[0].([]model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockBookMockRecorder) GetAllBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockBook)(nil).GetAllBooks))
}

// GetBookById mocks base method.
func (m *MockBook) GetBookById(id int) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", id)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockBookMockRecorder) GetBookById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockBook)(nil).GetBookById), id)
}

// UpdateBook mocks base method.
func (m *MockBook) UpdateBook(book model.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookMockRecorder) UpdateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBook)(nil).UpdateBook), book)
}
