package service

import (
	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data"
	"github.com/yunya101/ozon-task/internal/model"
)

type UserService struct {
	repo data.UserRepository
}

func NewUserService(repo data.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) AddUser(user *model.User) error {
	if err := s.repo.Insert(user); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("new user added")
	return nil
}

func (s *UserService) DeleteUserById(id int64) error {
	if err := s.repo.DeleteById(id); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("user was deleted")

	return nil
}
