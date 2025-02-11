package service_test

import (
	"testing"

	"github.com/yunya101/ozon-task/internal/data"
	inmem "github.com/yunya101/ozon-task/internal/data/inmemory"
	"github.com/yunya101/ozon-task/internal/model"
)

var postRepoInMem data.PostRepository = inmem.NewPostRepoInMem()
var userRepoInMem data.UserRepository = inmem.NewUserRepoInMem()

func TestAddPost(t *testing.T) {
	rightPost1 := &model.Post{
		Title:       "First",
		Text:        "blablalba",
		Author:      &model.User{ID: 1, Username: "Nikolay"},
		IsCommented: true,
	}

	rightPost2 := &model.Post{
		Title:  "Second",
		Text:   "blablabla",
		Author: &model.User{ID: 1, Username: "Nikolay"},
	}

	wrongPost1 := &model.Post{
		Title:       "",
		Text:        "",
		Author:      &model.User{ID: 2, Username: "Nika"},
		IsCommented: true,
	}

	wrongPost2 := &model.Post{
		Title:       "Helllo",
		Text:        "Wooorld",
		Author:      &model.User{ID: -1},
		IsCommented: false,
	}

	posts := []*model.Post{rightPost1, rightPost2, wrongPost1, wrongPost2}

	for i, p := range posts {
		err := postRepoInMem.Insert(p)

		if err != nil {
			if i == 2 {

			}
		}
	}

}

func TestLastest(t *testing.T) {

}

func addUsers() {
	u1 := &model.User{
		ID:       1,
		Username: "Nikolay",
	}

	u2 := &model.User{
		ID:       2,
		Username: "Nika",
	}

	u3 := &model.User{
		Username: "Sasha",
	}

	userRepoInMem.Insert(u1)
	userRepoInMem.Insert(u2)
	userRepoInMem.Insert(u3)
}
