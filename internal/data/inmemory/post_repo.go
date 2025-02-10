package data

import (
	apperrors "github.com/yunya101/ozon-task/internal/errors"
	"github.com/yunya101/ozon-task/internal/model"
)

type PostRepoInMem struct {
	posts   map[int64]*model.Post
	countId int64
}

func NewPostRepoInMem() *PostRepoInMem {
	return &PostRepoInMem{
		posts:   map[int64]*model.Post{},
		countId: 1,
	}
}

func (r *PostRepoInMem) Insert(post *model.Post) error {
	post.ID = r.countId

	post.Comments = make([]*model.Comment, 0)

	r.posts[post.ID] = post
	r.countId++

	return nil
}

func (r *PostRepoInMem) Lastest(page int) ([]*model.Post, error) {

	posts := make([]*model.Post, 0)
	page *= 10

	for i := r.countId - 1; i > 0; i-- {
		p := r.posts[i]

		posts = append(posts, p)
		page--

		if page <= 0 {
			return posts, nil
		}
	}

	return posts, nil
}

func (r *PostRepoInMem) GetById(id int64) (*model.Post, error) {
	p := r.posts[id]

	if p == nil {
		return nil, apperrors.ErrNotFound
	}

	return p, nil
}

func (r *PostRepoInMem) Update(post *model.Post) error {

	r.posts[post.ID] = post

	return nil

}
