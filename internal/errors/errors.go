package apperrors

import (
	"errors"

	"github.com/yunya101/ozon-task/internal/model"
)

var (
	ErrEmptyText     = errors.New("empty text")
	ErrLimitText     = errors.New("character count exceeded")
	ErrDoesntExist   = errors.New("data doesn't exist")
	ErrCannotComment = errors.New("now allowed to comment on this post")
	ErrNotFound      = errors.New("not found")
	ErrWrongName     = errors.New("wrong username")
	ErrWrongUserId   = errors.New("wrong user id")
)

func CheckPost(post *model.Post) error {

	if post.Text == "" {
		return ErrEmptyText
	}

	if len(post.Text) > 2000 {
		return ErrLimitText
	}

	if post.Title == "" || len(post.Title) < 3 {
		return ErrEmptyText
	}

	if post.Author.ID <= 0 {
		return ErrWrongUserId
	}

	return nil
}

func CheckComment(com *model.Comment) error {
	if com.Text == "" {
		return ErrEmptyText
	}

	if len(com.Text) > 500 {
		return ErrLimitText
	}

	return nil
}

func CheckUser(user *model.User) error {
	if len(user.Username) > 20 || len(user.Username) < 4 {
		return ErrWrongName
	}

	return nil
}
