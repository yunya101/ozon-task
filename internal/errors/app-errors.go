package apperrors

import (
	"errors"

	"github.com/yunya101/ozon-task/internal/model"
)

var (
	ErrEmptyText = errors.New("empty text")
	ErrLimitText = errors.New("character count exceeded")
)

func CheckPost(post *model.Post) error {

	if post.Text == "" {
		return ErrEmptyText
	}

	if len(post.Text) > 2000 {
		return ErrLimitText
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
