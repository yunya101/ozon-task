package service

import (
	"time"

	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type PostService struct {
	repo data.PostRepository
}

func NewPostService(r data.PostRepository) *PostService {
	return &PostService{
		repo: r,
	}
}

func (s *PostService) GetLastestPosts(page int) ([]*model.Post, error) {
	posts, err := s.repo.Lastest(page)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	config.InfoLog("getting lastest posts")

	return posts, nil
}

func (s *PostService) AddPost(post *model.Post) error {

	if err := apperrors.CheckPost(post); err != nil {
		config.ErrorLog(err)
		return err
	}

	post.LastCommentTime = time.Now()

	err := s.repo.Insert(post)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("post added")

	return err
}

func (s *PostService) GetPostById(id int64) (*model.Post, error) {

	post, err := s.repo.GetById(id)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	config.InfoLog("getting post by id")

	return post, nil
}
