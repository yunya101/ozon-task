package route_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yunya101/ozon-task/cmd/graph"
	"github.com/yunya101/ozon-task/cmd/route"
	"github.com/yunya101/ozon-task/internal/model"
)

// Mock для UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) AddUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Mock для PostService
type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) AddPost(post *model.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func setupRouter(userSvc *MockUserService, postSvc *MockPostService) *mux.Router {
	r := route.NewRouter(&graph.Resolver{}, userSvc, postSvc)
	r.SetRoutes()
	return r.GetMux()
}

func TestAddUser_Success(t *testing.T) {
	mockUserService := new(MockUserService)
	mockPostService := new(MockPostService)
	router := setupRouter(mockUserService, mockPostService)

	user := &model.User{Username: "yunya"}
	body, _ := json.Marshal(user)

	mockUserService.On("AddUser", user).Return(nil)

	req, _ := http.NewRequest("POST", "/user/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUserService.AssertExpectations(t)
}

func TestAddUser_InvalidJSON(t *testing.T) {
	mockUserService := new(MockUserService)
	mockPostService := new(MockPostService)
	router := setupRouter(mockUserService, mockPostService)

	req, _ := http.NewRequest("POST", "/user/add", bytes.NewBuffer([]byte("{fignya}")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockUserService.AssertNotCalled(t, "AddUser")
}

func TestAddUser_ServiceError(t *testing.T) {
	mockUserService := new(MockUserService)
	mockPostService := new(MockPostService)
	router := setupRouter(mockUserService, mockPostService)

	user := &model.User{Username: "nik"}
	body, _ := json.Marshal(user)

	mockUserService.On("AddUser", user).Return(errors.New("service error"))

	req, _ := http.NewRequest("POST", "/user/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockUserService.AssertExpectations(t)
}

func TestAddPost_Success(t *testing.T) {
	mockUserService := new(MockUserService)
	mockPostService := new(MockPostService)
	router := setupRouter(mockUserService, mockPostService)

	post := &model.Post{Title: "Fish", Text: "I like fish"}
	body, _ := json.Marshal(post)

	mockPostService.On("AddPost", post).Return(nil)

	req, _ := http.NewRequest("POST", "/post/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockPostService.AssertExpectations(t)
}

func TestAddPost_InvalidJSON(t *testing.T) {
	mockUserService := new(MockUserService)
	mockPostService := new(MockPostService)
	router := setupRouter(mockUserService, mockPostService)

	req, _ := http.NewRequest("POST", "/post/add", bytes.NewBuffer([]byte("{dich}")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockPostService.AssertNotCalled(t, "AddPost")
}

func TestAddPost_ServiceError(t *testing.T) {
	mockUserService := new(MockUserService)
	mockPostService := new(MockPostService)
	router := setupRouter(mockUserService, mockPostService)

	post := &model.Post{Title: "Test", Text: "testing"}
	body, _ := json.Marshal(post)

	mockPostService.On("AddPost", post).Return(errors.New("service error"))

	req, _ := http.NewRequest("POST", "/post/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockPostService.AssertExpectations(t)
}
