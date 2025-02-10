package graph

import (
	"sync"

	"github.com/yunya101/ozon-task/internal/model"
	"github.com/yunya101/ozon-task/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	mu             *sync.Mutex
	subs           map[int64]chan *model.Comment
	postService    *service.PostService
	userService    *service.UserService
	commentService *service.CommsService
}

func NewResolver(pServ *service.PostService, uServ *service.UserService, cServ *service.CommsService) *Resolver {

	return &Resolver{
		postService:    pServ,
		userService:    uServ,
		commentService: cServ,
		mu:             &sync.Mutex{},
		subs:           make(map[int64]chan *model.Comment),
	}

}

func (r *Resolver) SetPostService(s *service.PostService) {
	r.postService = s
}

func (r *Resolver) SetUserService(s *service.UserService) {
	r.userService = s
}

func (r *Resolver) SetCommentService(s *service.CommsService) {
	r.commentService = s
}
