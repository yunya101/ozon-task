package testhelper

import (
	"time"

	"github.com/yunya101/ozon-task/internal/model"
)

func GetNewPosts() []*model.Post {
	posts := make([]*model.Post, 0)
	c1 := &model.Comment{
		ID:        1,
		Text:      "First comment!",
		Author:    &model.User{ID: 1, Username: "Nikolay"},
		PostID:    1,
		ParentID:  0,
		CreatedAt: time.Now(),
	}

	c2 := &model.Comment{
		ID:        2,
		Text:      "2 comment!",
		Author:    &model.User{ID: 2, Username: "Nika"},
		PostID:    1,
		ParentID:  1,
		CreatedAt: time.Now(),
	}

	c3 := &model.Comment{
		ID:        3,
		Text:      "3 comment!",
		Author:    &model.User{ID: 1, Username: "Nikolay"},
		PostID:    1,
		ParentID:  2,
		CreatedAt: time.Now(),
	}

	c4 := &model.Comment{
		ID:        4,
		Text:      "4 comment!",
		Author:    &model.User{ID: 3, Username: "Vika"},
		PostID:    1,
		ParentID:  3,
		CreatedAt: time.Now(),
	}

	c1.Comments = []*model.Comment{c2}
	c2.Comments = []*model.Comment{c3}
	c3.Comments = []*model.Comment{c4}

	c11 := &model.Comment{
		ID:        5,
		Text:      "5 comment!",
		Author:    &model.User{ID: 2, Username: "Nika"},
		PostID:    2,
		ParentID:  0,
		CreatedAt: time.Now(),
	}
	c12 := &model.Comment{
		ID:        6,
		Text:      "6 comment!",
		Author:    &model.User{ID: 1, Username: "Nikolay"},
		PostID:    2,
		ParentID:  0,
		CreatedAt: time.Now(),
	}

	p1 := &model.Post{
		ID:          1,
		Title:       "First!",
		Text:        "Just a normal post",
		Author:      &model.User{ID: 1, Username: "Nikolay"},
		IsCommented: true,
		Comments:    []*model.Comment{c1},
	}

	p2 := &model.Post{
		ID:          2,
		Title:       "Second!",
		Text:        "Just a normal post",
		Author:      &model.User{ID: 2, Username: "Nika"},
		IsCommented: false,
	}

	p3 := &model.Post{
		ID:          3,
		Title:       "Третий!",
		Text:        "Просто пост!",
		Author:      &model.User{ID: 3, Username: "Vika"},
		IsCommented: true,
		Comments:    []*model.Comment{c11, c12},
	}

	posts = append(posts, p1, p2, p3)
	return posts
}
