package data

import (
	"database/sql"

	"github.com/yunya101/ozon-task/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) SetDB(db *sql.DB) {
	r.db = db
}

func (r *UserRepository) InsertUser(user *model.User) error {
	stmt := `INSERT INTO users (username)
	VALUES ($1)`

	_, err := r.db.Exec(stmt, user.Username)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	stmt := `UPDATE users SET username = $1 WHERE id = $2`

	_, err := r.db.Exec(stmt, user.Username, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUserById(id int64) error {
	stmt := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(stmt, id)

	if err != nil {
		return err
	}

	return nil
}
