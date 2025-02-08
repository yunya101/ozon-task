package model

import "time"

type Comment struct {
	ID        int64      `json:"id"`
	Author    User       `json:"author"`
	Text      string     `json:"text"`
	PostID    int64      `json:"postId"`
	Comments  []*Comment `json:"Comments"`
	ParentID  int64      `json:"parentId"`
	CreatedAt time.Time  `json:"createAt"`
}
