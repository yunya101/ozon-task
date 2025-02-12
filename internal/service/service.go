package service

import "github.com/yunya101/ozon-task/internal/model"

type UserServiceInt interface {
	AddUser(user *model.User) error
}

type PostServiceInt interface {
	AddPost(post *model.Post) error
}
