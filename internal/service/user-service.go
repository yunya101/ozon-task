package service

import (
	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data/postgres"
	"github.com/yunya101/ozon-task/internal/model"
)

type UserService struct {
	repo *data.UserRepo
}

func (s *UserService) SetRepo(r *data.UserRepo) {
	s.repo = r
}

func (s *UserService) AddUser(user *model.User) error {
	if err := s.repo.InsertUser(user); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("new user added")
	return nil
}

func (s *UserService) UpdateUser(user *model.User) error {
	if err := s.repo.UpdateUser(user); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("user was updated")

	return nil
}

func (s *UserService) DeleteUserById(id int64) error {
	if err := s.repo.DeleteUserById(id); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("user was deleted")

	return nil
}
