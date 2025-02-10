package data

import (
	"database/sql"

	"github.com/yunya101/ozon-task/internal/config"
	"github.com/yunya101/ozon-task/internal/model"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) Insert(com *model.Comment) error {
	stmt := `INSERT INTO comments (author, text, post, parent, createAt)
	VALUES ($1, $2, $3, $4, $5)`

	var err error

	if com.ParentID > 0 {
		_, err = r.db.Exec(stmt, com.Author.ID, com.Text, com.PostID, com.ParentID, com.CreatedAt)
	} else {
		_, err = r.db.Exec(stmt, com.Author.ID, com.Text, com.PostID, nil, com.CreatedAt)
	}

	if err != nil {
		config.ErrorLog(err)
		return err
	}

	return nil
}

func (r *CommentRepo) Update(com *model.Comment) error {
	stmt := `UPDATE comments
	SET text = $1 WHERE id = $2`

	_, err := r.db.Exec(stmt, com.Text, com.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepo) DeleteCommentById(id int64) error {
	stmt := `DELETE FROM comments WHERE id = $1`

	_, err := r.db.Exec(stmt, id)

	if err != nil {
		return err
	}

	return nil
}
