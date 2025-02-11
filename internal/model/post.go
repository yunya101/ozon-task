package model

type Post struct {
	ID          int64      `json:"id"`
	Author      *User      `json:"author"`
	Title       string     `json:"title"`
	Text        string     `json:"text"`
	Comments    []*Comment `json:"comments"`
	IsCommented bool       `json:"isCommented"`
}
