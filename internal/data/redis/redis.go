package data

import (
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yunya101/ozon-task/internal/config"
	pg "github.com/yunya101/ozon-task/internal/data/postgres"
	"github.com/yunya101/ozon-task/internal/model"
)

type RedisRepo struct {
	repo *redis.Client
	pg   *pg.PostRepo
}

func NewRedisRepo(c *redis.Client, pg *pg.PostRepo) *RedisRepo {
	return &RedisRepo{
		repo: c,
		pg:   pg,
	}
}

func (r *RedisRepo) AddPost(post *model.Post) error {
	key := fmt.Sprintf("%v", post.ID)

	json, err := json.Marshal(post)

	if err != nil {
		return err
	}

	config.InfoLog("adding post to redis")

	return r.repo.Set(config.Ctx, key, json, 0).Err()
}

func (r *RedisRepo) DeletePost(id int64) error {

	key := fmt.Sprintf("%v", id)

	config.InfoLog("removing post from redis")

	return r.repo.Del(config.Ctx, key).Err()
}

func (r *RedisRepo) GetPostById(id int64) (*model.Post, bool, error) {

	key := fmt.Sprintf("%v", id)

	exist, err := r.repo.Exists(config.Ctx, key).Result()

	if err != nil {
		config.ErrorLog(err)
		return nil, false, err
	}

	if exist > 0 {
		bytePost, err := r.repo.Get(config.Ctx, key).Bytes()

		if err != nil {
			config.ErrorLog(err)
			return nil, false, err
		}

		post := &model.Post{}

		if err := json.Unmarshal(bytePost, post); err != nil {
			config.ErrorLog(err)
			return nil, false, err
		}
		config.InfoLog("getting post by id from redis")
		return post, true, nil
	}

	return nil, false, nil
}

func (r *RedisRepo) GetAllPosts() ([]*model.Post, error) {

	keys, err := r.repo.Keys(config.Ctx, "*").Result()

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	posts := make([]*model.Post, 0)

	for _, k := range keys {
		bytePost, err := r.repo.Get(config.Ctx, k).Bytes()

		if err != nil {
			return nil, err
		}

		post := &model.Post{}
		err = json.Unmarshal(bytePost, post)

		if err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		posts = append(posts, post)
	}

	config.InfoLog("getting all posts from redis")

	return posts, nil
}
