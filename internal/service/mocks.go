package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/yunya101/ozon-task/internal/model"
)

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
