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
type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) Lastest(page int) ([]*model.Post, error) {
	args := m.Called(page)
	return args.Get(0).([]*model.Post), args.Error(1)
}

func (m *MockPostRepository) Insert(post *model.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostRepository) GetById(id int64) (*model.Post, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Post), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPostRepository) Update(post *model.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func TestPostService_GetLastestPosts_Success(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := service.NewPostService(mockRepo)

	expectedPosts := []*model.Post{
		{ID: 1, Title: "Post 1", Text: "Content 1"},
		{ID: 2, Title: "Post 2", Text: "Content 2"},
	}

	mockRepo.On("Lastest", 1).Return(expectedPosts, nil)

	posts, err := service.GetLastestPosts(1)

	assert.NoError(t, err)
	assert.Len(t, posts, 2)
	mockRepo.AssertExpectations(t)
}

func TestPostService_GetLastestPosts_Fail(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := service.NewPostService(mockRepo)

	mockRepo.On("Lastest", 1).Return([]*model.Post{}, errors.New("db error"))

	posts, err := service.GetLastestPosts(1)

	assert.Error(t, err)
	assert.Nil(t, posts)
	mockRepo.AssertExpectations(t)
}

func TestPostService_AddPost_Success(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := service.NewPostService(mockRepo)

	post := &model.Post{Title: "  New Post  ", Text: "  Some content  ", Author: &model.User{ID: 1}, IsCommented: true}
	mockRepo.On("Insert", mock.Anything).Return(nil)

	err := service.AddPost(post)

	assert.NoError(t, err)
	assert.Equal(t, "New Post", post.Title)
	assert.Equal(t, "Some content", post.Text)
	mockRepo.AssertExpectations(t)
}

func TestPostService_AddPost_Fail(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := service.NewPostService(mockRepo)

	post := &model.Post{Title: "", Text: ""}
	mockRepo.On("Insert", post).Return(apperrors.ErrEmptyText)

	err := service.AddPost(post)

	assert.Error(t, err)
	assert.Equal(t, apperrors.ErrEmptyText, err)
}

func TestPostService_GetPostById_Success(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := service.NewPostService(mockRepo)

	expectedPost := &model.Post{ID: 1, Title: "Found Post", Text: "Some content"}
	mockRepo.On("GetById", int64(1)).Return(expectedPost, nil)

	post, err := service.GetPostById(1)

	assert.NoError(t, err)
	assert.Equal(t, "Found Post", post.Title)
	mockRepo.AssertExpectations(t)
}

func TestPostService_GetPostById_Fail(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := service.NewPostService(mockRepo)

	mockRepo.On("GetById", int64(1)).Return(nil, errors.New("not found"))

	post, err := service.GetPostById(1)

	assert.Error(t, err)
	assert.Nil(t, post)
	assert.Equal(t, "not found", err.Error())
	mockRepo.AssertExpectations(t)
}
