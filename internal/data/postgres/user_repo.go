package data

import (
	"database/sql"

	"github.com/yunya101/ozon-task/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Insert(user *model.User) error {
	stmt := `INSERT INTO users (username)
	VALUES ($1)`

	_, err := r.db.Exec(stmt, user.Username)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Update(user *model.User) error {
	stmt := `UPDATE users SET username = $1 WHERE id = $2`

	_, err := r.db.Exec(stmt, user.Username, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) DeleteById(id int64) error {
	stmt := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(stmt, id)

	if err != nil {
		return err
	}

	return nil
}
