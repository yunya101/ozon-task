package repository

import (
	"github.com/yunya101/ozon-task/internal/model"
)

type UserRepository interface {
	InsertUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUserById(id int64) error
}

type PostRepository interface {
	GetSubsPostsByUserId(id int64) ([]*model.Post, error)
	GetLastestPosts(page int) ([]*model.Post, error)
	InsertPost(post *model.Post) error
	GetPostById(id int64) (*model.Post, error)
	AddUserInPost(postId int64, userID int64) error
	UpdatePost(post *model.Post) error
	DeletePostById(id int64) error
}

type CommentRepository interface {
	InsertComment(com *model.Comment) error
	UpdateComment(com *model.Comment) error
	DeleteCommentById(id int64) error
}
