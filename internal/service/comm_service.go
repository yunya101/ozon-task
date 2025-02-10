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

func (s *CommsService) AddComment(com *model.Comment) error {
	if err := apperrors.CheckComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	com.CreatedAt = time.Now()

	post, err := s.postRepo.GetById(com.PostID)
	if err != nil {
		return err
	}

	if !post.IsCommented {
		return apperrors.ErrCannotComment
	}

	com.Comments = make([]*model.Comment, 0)
	if err := s.repo.Insert(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	if com.ParentID <= 0 {
		post.Comments = append(post.Comments, com)
	}

	post.LastCommentTime = com.CreatedAt

	if err := s.postRepo.Update(post); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("new comment added")

	return nil
}
