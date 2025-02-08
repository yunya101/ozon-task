package service

import (
	"github.com/yunya101/ozon-task/internal/config"
	repository "github.com/yunya101/ozon-task/internal/data"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type PostService struct {
	repo repository.PostRepository
}

func (s *PostService) SetRepo(r repository.PostRepository) {
	s.repo = r
}

func (s *PostService) GetLastestPosts(page int) ([]*model.Post, error) {
	posts, err := s.repo.GetLastestPosts(page)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetUsersSubsPosts(id int64) ([]*model.Post, error) {
	posts, err := s.repo.GetSubsPostsByUserId(id)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	return posts, nil
}

func (s *PostService) AddPost(post *model.Post) error {

	if err := apperrors.CheckPost(post); err != nil {
		config.ErrorLog(err)
		return err
	}

	err := s.repo.InsertPost(post)

	if err != nil {
		config.ErrorLog(err)
	}

	return err
}

func (s *PostService) GetPostById(id int64) (*model.Post, error) {
	post, err := s.repo.GetPostById(id)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	return post, nil
}

func (s *PostService) SupscribeUser(postID int64, userID int64) error {
	err := s.repo.AddUserInPost(postID, userID)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (s *PostService) UpdatePost(post *model.Post) error {

	if err := apperrors.CheckPost(post); err != nil {
		config.ErrorLog(err)
		return err
	}

	err := s.repo.UpdatePost(post)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (s *PostService) DeletePost(id int64) error {

	if err := s.repo.DeletePostById(id); err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}
