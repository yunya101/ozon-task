package data

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/yunya101/ozon-task/internal/config"
	"github.com/yunya101/ozon-task/internal/model"
	"github.com/yunya101/ozon-task/pkg/lib"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) Lastest(page int) ([]*model.Post, error) {

	stmt := `SELECT p.*, u.username FROM posts p
	JOIN users u ON p.author = u.id
	LIMIT 10
	OFFSET $1`

	posts := make([]*model.Post, 0)

	rows, err := r.db.Query(stmt, page)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		p := &model.Post{}
		u := &model.User{}

		if err = rows.Scan(&p.ID, &u.ID, &p.Title, &p.Text, &p.IsCommented, &p.CountComms, &p.LastCommentTime, &u.Username); err != nil {
			return nil, err
		}

		p.Author = u
		p.Comments = make([]*model.Comment, 0)
		posts = append(posts, p)

	}

	posts, err = r.getCommentsForPosts(posts)
	if err != nil {
		return nil, err
	}

	return posts, nil

}

func (r *PostRepo) getCommentsForPosts(posts []*model.Post) ([]*model.Post, error) {

	stmt := `select c.id, c.author, c.text, c.post, c.parent, c.createat, u.username from comments c
	JOIN users u ON c.author = u.id
	WHERE c.post = ANY($1)`

	postIds := make([]int64, 0)
	for _, p := range posts {
		postIds = append(postIds, p.ID)
	}

	rows, err := r.db.Query(stmt, pq.Array(postIds))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := make([]*model.Comment, 0)
	childComments := make(map[int64][]*model.Comment)

	for rows.Next() {
		c := &model.Comment{}
		u := &model.User{}

		var parentNull sql.NullInt64

		if err = rows.Scan(&c.ID, &u.ID, &c.Text, &c.PostID, &parentNull, &c.CreatedAt, &u.Username); err != nil {
			return nil, err
		}

		if parentNull.Valid {
			c.ParentID = parentNull.Int64
		}

		c.Author = u

		comments = append(comments, c)

		// TODO - решить проблему вложенности комментариев

		if childComments[c.ParentID] == nil && c.ParentID > 0 {
			childComments[c.ParentID] = make([]*model.Comment, 0)
		}

		childComments[c.ParentID] = append(childComments[c.ParentID], c)
	}

	for i, c := range comments {
		com, exist := childComments[c.ID]
		if exist {
			comments[i].Comments = com
			delete(childComments, c.ID)
		}
	}

	for j, p := range posts {
		for i := 0; i < len(comments); i++ {
			if comments[i].PostID == p.ID && comments[i].ParentID > 0 {
				posts[j].Comments = append(posts[j].Comments, comments[i])
				comments = lib.RemoveCommentFromSlice(comments, i)
				i--
			}
		}
	}

	return posts, nil

}

func (r *PostRepo) GetById(id int64) (*model.Post, error) {

	stmt := `SELECT p.*, u.username FROM posts p
	JOIN users u ON p.author = u.id
	WHERE p.id = $1`

	row := r.db.QueryRow(stmt, id)

	p := &model.Post{}
	u := &model.User{}

	if err := row.Scan(&p.ID, &u.ID, &p.Title, &p.Text, &p.IsCommented, &p.CountComms, &p.LastCommentTime, &u.Username); err != nil {
		return nil, err
	}

	p.Author = u

	stub := []*model.Post{p}

	stub, err := r.getCommentsForPosts(stub)
	if err != nil {
		return nil, err
	}

	return stub[0], nil
}

func (r *PostRepo) Insert(post *model.Post) error {
	stmt := `INSERT INTO posts (author, title, text, iscommented, countcomments, lastcommenttime)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(stmt, post.Author.ID, post.Title, post.Text, post.IsCommented, post.CountComms, post.LastCommentTime)
	if err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (r *PostRepo) Update(post *model.Post) error {
	stmt := `UPDATE posts
	SET title = $1, text = $2, iscommented = $3, countcomments = $4, lastcommenttime = $5`

	_, err := r.db.Exec(stmt, post.Title, post.Text, post.IsCommented, post.CountComms, post.LastCommentTime)
	if err != nil {
		return err
	}

	return nil
}
