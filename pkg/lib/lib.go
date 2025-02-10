package lib

import "github.com/yunya101/ozon-task/internal/model"

func RemoveCommentFromSlice(slice []*model.Comment, i int) []*model.Comment {
	return append(slice[:i], slice[i+1:]...)
}
