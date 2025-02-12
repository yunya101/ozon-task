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

// Mock для CommentRepository
type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) Insert(com *model.Comment) (int64, error) {
	args := m.Called(com)
	return args.Get(0).(int64), args.Error(1)
}

func TestCommsService_AddComment_Success(t *testing.T) {
	mockCommentRepo := new(MockCommentRepository)
	mockPostRepo := new(MockPostRepository)
	service := service.NewCommService(mockCommentRepo, mockPostRepo)

	comment := &model.Comment{
		Text:   "Nice post!",
		PostID: 1,
	}

	post := &model.Post{
		ID:          1,
		IsCommented: true,
		Comments:    []*model.Comment{},
	}

	mockPostRepo.On("GetById", int64(1)).Return(post, nil)
	mockCommentRepo.On("Insert", mock.MatchedBy(func(c *model.Comment) bool {
		return c.Text == "Nice post!" && c.PostID == 1
	})).Return(int64(1), nil)

	id, err := service.AddComment(comment)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
	mockPostRepo.AssertExpectations(t)
	mockCommentRepo.AssertExpectations(t)
}

func TestCommsService_AddComment_Invalid(t *testing.T) {
	mockCommentRepo := new(MockCommentRepository)
	mockPostRepo := new(MockPostRepository)
	service := service.NewCommService(mockCommentRepo, mockPostRepo)

	comment := &model.Comment{}

	id, err := service.AddComment(comment)

	assert.Error(t, err)
	assert.Equal(t, int64(0), id)
	mockPostRepo.AssertNotCalled(t, "GetById")
	mockCommentRepo.AssertNotCalled(t, "Insert")
}

func TestCommsService_AddComment_PostNotFound(t *testing.T) {
	mockCommentRepo := new(MockCommentRepository)
	mockPostRepo := new(MockPostRepository)
	service := service.NewCommService(mockCommentRepo, mockPostRepo)

	comment := &model.Comment{
		Text:   "Great post!",
		PostID: 99,
	}

	mockPostRepo.On("GetById", int64(99)).Return(nil, errors.New("post not found"))

	id, err := service.AddComment(comment)

	assert.Error(t, err)
	assert.Equal(t, "post not found", err.Error())
	assert.Equal(t, int64(0), id)
	mockPostRepo.AssertExpectations(t)
	mockCommentRepo.AssertNotCalled(t, "Insert")
}

func TestCommsService_AddComment_CommentsNotAllowed(t *testing.T) {
	mockCommentRepo := new(MockCommentRepository)
	mockPostRepo := new(MockPostRepository)
	service := service.NewCommService(mockCommentRepo, mockPostRepo)

	comment := &model.Comment{
		Text:   "Interesting!",
		PostID: 2,
	}

	post := &model.Post{
		ID:          2,
		IsCommented: false,
	}

	mockPostRepo.On("GetById", int64(2)).Return(post, nil)

	id, err := service.AddComment(comment)

	assert.Error(t, err)
	assert.Equal(t, apperrors.ErrCannotComment, err)
	assert.Equal(t, int64(0), id)
	mockPostRepo.AssertExpectations(t)
	mockCommentRepo.AssertNotCalled(t, "Insert")
}

func TestCommsService_AddComment_InsertFail(t *testing.T) {
	mockCommentRepo := new(MockCommentRepository)
	mockPostRepo := new(MockPostRepository)
	service := service.NewCommService(mockCommentRepo, mockPostRepo)

	comment := &model.Comment{
		Text:   "Very informative!",
		PostID: 3,
	}

	post := &model.Post{
		ID:          3,
		IsCommented: true,
		Comments:    []*model.Comment{},
	}

	mockPostRepo.On("GetById", int64(3)).Return(post, nil)
	mockCommentRepo.On("Insert", mock.Anything).Return(int64(0), errors.New("insert error"))

	id, err := service.AddComment(comment)

	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())
	assert.Equal(t, int64(0), id)
	mockPostRepo.AssertExpectations(t)
	mockCommentRepo.AssertExpectations(t)
}
