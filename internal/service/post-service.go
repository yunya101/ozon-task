package service

import (
	"time"

	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data/postgres"
	redis "github.com/yunya101/ozon-task/internal/data/redis"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type PostService struct {
	repo  *data.PostRepo
	redis *redis.RedisRepo
}

func NewPostService(repo *data.PostRepo, redis *redis.RedisRepo) *PostService {
	return &PostService{
		repo:  repo,
		redis: redis,
	}
}

func (s *PostService) GetLastestPosts(page int) ([]*model.Post, error) {
	posts, err := s.repo.GetLastestPosts(page)

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

	err := s.repo.InsertPost(post)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("post added")

	return err
}

func (s *PostService) GetPostById(id int64) (*model.Post, error) {

	post, exist, err := s.redis.GetPostById(id)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	if exist {
		return post, nil
	}

	post, err = s.repo.GetPostById(id)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	config.InfoLog("getting post by id")

	return post, nil
}
