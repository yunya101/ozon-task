package service

import (
	"github.com/yunya101/ozon-task/internal/config"
	repository "github.com/yunya101/ozon-task/internal/data"
	"github.com/yunya101/ozon-task/internal/model"
)

type UserService struct {
	repo repository.UserRepository
}

func (s *UserService) SetRepo(r repository.UserRepository) {
	s.repo = r
}

func (s *UserService) AddUser(user *model.User) error {
	if err := s.repo.InsertUser(user); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (s *UserService) UpdateUser(user *model.User) error {
	if err := s.repo.UpdateUser(user); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (s *UserService) DeleteUserById(id int64) error {
	if err := s.repo.DeleteUserById(id); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}
