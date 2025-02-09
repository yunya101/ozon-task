package model

import "time"

type Comment struct {
	ID        int64      `json:"id"`
	Author    int64      `json:"author"`
	Text      string     `json:"text"`
	PostID    int64      `json:"post"`
	Comments  []*Comment `json:"Comments"`
	ParentID  int64      `json:"parent"`
	CreatedAt time.Time  `json:"createAt"`
}
