package model

import "time"

type Post struct {
	ID              int64      `json:"id"`
	Author          int64      `json:"author"`
	Title           string     `json:"title"`
	Text            string     `json:"text"`
	Subs            []*User    `json:"subs"`
	Comments        []*Comment `json:"comments"`
	CountComms      int        `json:"countComments"`
	IsCommented     bool       `json:"isCommented"`
	LastCommentTime time.Time
	Popularity      float64
}

func (p *Post) CalcPopularity() {
	p.Popularity = float64(p.CountComms) * float64(time.Since(p.LastCommentTime).Hours()+1)
}
