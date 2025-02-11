package service

import (
	"time"

	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type CommsService struct {
	repo     data.CommentRepository
	postRepo data.PostRepository
}

func NewCommService(r data.CommentRepository, postRepo data.PostRepository) *CommsService {
	return &CommsService{
		repo:     r,
		postRepo: postRepo,
	}
}

func (s *CommsService) AddComment(com *model.Comment) (int64, error) {
	if err := apperrors.CheckComment(com); err != nil {
		config.ErrorLog(err)
		return 0, err
	}

	com.CreatedAt = time.Now()

	post, err := s.postRepo.GetById(com.PostID)
	if err != nil {
		return 0, err
	}

	if !post.IsCommented {
		return 0, apperrors.ErrCannotComment
	}

	com.Comments = make([]*model.Comment, 0)
	id, err := s.repo.Insert(com)
	if err != nil {
		return 0, err
	}

	if com.ParentID <= 0 {
		post.Comments = append(post.Comments, com)
	}

	post.LastCommentTime = com.CreatedAt

	if err := s.postRepo.Update(post); err != nil {
		config.ErrorLog(err)
		return 0, err
	}

	config.InfoLog("new comment added")

	return id, nil
}
