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

func (s *PostService) SetRepo(r *data.PostRepo) {
	s.repo = r
}

func (s *PostService) SetRedis(redis *redis.RedisRepo) {
	s.redis = redis
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

func (s *PostService) GetSubsPostsByUserId(id int64) ([]*model.Post, error) {

	postsIds, err := s.repo.GetPostsIdsByUserId(id)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	posts := make([]*model.Post, 0)

	for _, pId := range postsIds {

		post, exist, err := s.redis.GetPostById(pId)

		if err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		if exist {
			posts = append(posts, post)
			config.InfoLog("getting post from redis")
		} else {
			post, err = s.repo.GetPostById(pId)
			config.InfoLog("getting post from postgres")
			if err != nil {
				config.ErrorLog(err)
				return nil, err
			}
		}

		posts = append(posts, post)

	}

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

func (s *PostService) SupscribeUser(postID int64, userID int64) error {
	err := s.repo.AddUserInPost(postID, userID)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	config.InfoLog("new user subscribe")

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

	_, exist, err := s.redis.GetPostById(post.ID)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	if exist {
		if err = s.redis.AddPost(post); err != nil {
			return err
		}
	}

	config.InfoLog("post was updated")
	return nil
}

func (s *PostService) DeletePost(id int64) error {

	if err := s.repo.DeletePostById(id); err != nil {
		config.ErrorLog(err)
		return err
	}

	_, exist, err := s.redis.GetPostById(id)

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	if exist {
		s.redis.DeletePost(id)
	}

	config.InfoLog("post was deleted")

	return nil
}
