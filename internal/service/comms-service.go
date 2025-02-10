package service

import (
	"time"

	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data/postgres"
	redis "github.com/yunya101/ozon-task/internal/data/redis"
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type CommsService struct {
	repo     *data.CommentRepo
	redis    *redis.RedisRepo
	postRepo *data.PostRepo
}

func NewCommService(r *data.CommentRepo, redis *redis.RedisRepo, postRepo *data.PostRepo) *CommsService {
	return &CommsService{
		repo:     r,
		redis:    redis,
		postRepo: postRepo,
	}
}

func (s *CommsService) AddComment(com *model.Comment) error {
	if err := apperrors.CheckComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	com.CreatedAt = time.Now()

	post, exist, err := s.redis.GetPostById(com.PostID)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	if !exist {
		post, err = s.postRepo.GetPostById(com.PostID)

		if err != nil {
			config.ErrorLog(err)
			return err
		}
	}

	post.LastCommentTime = com.CreatedAt
	post.CountComms++
	post.CalcPopularity()

	if post.Popularity > config.PopularityThreshold {
		s.redis.AddPost(post)
	} else {
		s.redis.DeletePost(post.ID)
	}

	if err := s.repo.InsertComment(com); err != nil {
		config.ErrorLog(err)
		return err
	}

	if err := s.postRepo.UpdatePost(post); err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("new comment")

	return nil
}
