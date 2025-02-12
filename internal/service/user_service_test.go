package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
	"github.com/yunya101/ozon-task/internal/service"
)

// Мок репозитория
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Insert(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteById(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserService_AddUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.NewUserService(mockRepo)

	user := &model.User{Username: "   test_user   "}
	mockRepo.On("Insert", mock.Anything).Return(nil)

	err := service.AddUser(user)

	assert.NoError(t, err)
	assert.Equal(t, "test_user", user.Username)
	mockRepo.AssertCalled(t, "Insert", user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_AddUser_Fail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.NewUserService(mockRepo)

	user := &model.User{Username: "invalid_user"}
	mockRepo.On("Insert", user).Return(errors.New("db error"))

	err := service.AddUser(user)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	mockRepo.AssertCalled(t, "Insert", user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUserById_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.NewUserService(mockRepo)

	mockRepo.On("DeleteById", int64(1)).Return(nil)

	err := service.DeleteUserById(1)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteById", int64(1))
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUserById_Fail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.NewUserService(mockRepo)

	mockRepo.On("DeleteById", int64(1)).Return(apperrors.ErrDoesntExist)

	err := service.DeleteUserById(1)

	assert.Error(t, err)
	assert.Equal(t, apperrors.ErrDoesntExist, err)
	mockRepo.AssertExpectations(t)
}
