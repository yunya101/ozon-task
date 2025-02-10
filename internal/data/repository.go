package data

import "github.com/yunya101/ozon-task/internal/model"

type PostRepository interface {
	Insert(*model.Post) error
	Lastest(int) ([]*model.Post, error)
	GetById(int64) (*model.Post, error)
	Update(*model.Post) error
}

type UserRepository interface {
	Insert(*model.User) error
	DeleteById(int64) error
}

type CommentRepository interface {
	Insert(*model.Comment) error
}
