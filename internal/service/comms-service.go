package service

import (
	"github.com/yunya101/ozon-task/internal/config"
	repository "github.com/yunya101/ozon-task/internal/data"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type CommsService struct {
	repo repository.CommentRepository
}

func (s *CommsService) SetRepo(r repository.CommentRepository) {
	s.repo = r
}

func (s *CommsService) AddComment(com *model.Comment) error {
	if err := apperrors.CheckComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	if err := s.repo.InsertComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (s *CommsService) UpdateComment(com *model.Comment) error {
	if err := apperrors.CheckComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	if err := s.repo.UpdateComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (s *CommsService) DeleteComment(id int64) error {
	if err := s.repo.DeleteCommentById(id); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}
