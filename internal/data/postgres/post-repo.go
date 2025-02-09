package data

import (
	"database/sql"

	"github.com/yunya101/ozon-task/internal/config"
	"github.com/yunya101/ozon-task/internal/model"
)

type PostRepo struct {
	db *sql.DB
}

func (r *PostRepo) SetDB(db *sql.DB) {
	r.db = db
}

func (r *PostRepo) GetPostsIdsByUserId(id int64) ([]int64, error) {
	stmt := `SELECT post FROM users_posts WHERE user = $1`

	rows, err := r.db.Query(stmt, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	postsIds := make([]int64, 0)

	for rows.Next() {
		var postId int64
		if err := rows.Scan(&postId); err != nil {
			return nil, err
		}
		postsIds = append(postsIds, postId)
	}

	return postsIds, nil

}

func (r *PostRepo) GetSubsPostsByUserId(id int64) ([]*model.Post, error) {

	stmt := `SELECT * FROM posts p
	JOIN users_posts up ON p.id = up.post
	WHERE up.user_id = $1`

	fRows, err := r.db.Query(stmt, id)

	if err != nil {
		return nil, err
	}

	defer fRows.Close()

	posts := make([]*model.Post, 0)

	for fRows.Next() {
		p := &model.Post{}

		if err := fRows.Scan(&p.ID, &p.Author, &p.Title, &p.Text, &p.IsCommented, &p.CountComms); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	for i, p := range posts {
		p, err := r.getAllUsersFromPost(p)

		if err != nil {
			return nil, err
		}

		if p.IsCommented {
			p, err = r.getAllCommsFromPost(p)
			if err != nil {
				return nil, err
			}
		}
		posts[i] = p
	}

	return posts, nil

}

func (r *PostRepo) getAllUsersFromPost(post *model.Post) (*model.Post, error) {
	stmt := `SELECT u.* FROM users u
	JOIN users_posts up ON up.user_id = u.id
	WHERE up.post = $1`

	rows, err := r.db.Query(stmt, post.ID)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	defer rows.Close()

	users := make([]*model.User, 0)

	for rows.Next() {
		u := &model.User{}

		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		users = append(users, u)
	}

	post.Subs = users

	return post, nil
}

func (r *PostRepo) getAllCommsFromPost(post *model.Post) (*model.Post, error) {
	stmt := `SELECT * FROM comments WHERE post = $1`

	rows, err := r.db.Query(stmt, post.ID)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	defer rows.Close()

	comms := make([]*model.Comment, 0)

	for rows.Next() {
		c := &model.Comment{}
		var parent sql.NullInt64

		if err := rows.Scan(&c.ID, &c.Author, &c.Text, &c.PostID, &parent, &c.CreatedAt); err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		if parent.Valid {
			c.ParentID = parent.Int64
		}
		comms = append(comms, c)
	}

	for i, c := range comms {
		c, err := r.getAllCommsFromComm(c)

		if err != nil {
			config.ErrorLog(err)
			return nil, err
		}
		comms[i] = c
	}

	post.Comments = comms
	return post, nil
}

func (r *PostRepo) getAllCommsFromComm(comm *model.Comment) (*model.Comment, error) {
	stmt := `SELECT * FROM comments
	WHERE parent = $1`

	rows, err := r.db.Query(stmt, comm.ID)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	defer rows.Close()

	comms := make([]*model.Comment, 0)

	for rows.Next() {
		c := &model.Comment{}

		var parent sql.NullInt64
		if err := rows.Scan(&c.ID, &c.Author, &c.Text, &c.PostID, &parent, &c.CreatedAt); err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		if parent.Valid {
			c.ParentID = parent.Int64
		}

		c, err = r.getAllCommsFromComm(c)

		if err != nil {
			return nil, err
		}

		comms = append(comms, c)
	}

	comm.Comments = comms

	return comm, nil
}

func (r *PostRepo) GetLastestPosts(page int) ([]*model.Post, error) {
	stmt := `SELECT * FROM posts
	LIMIT 10 OFFSET $1 * 10`

	rows, err := r.db.Query(stmt, page)

	if err != nil {
		config.ErrorLog(err)
		return nil, err
	}

	defer rows.Close()

	posts := make([]*model.Post, 0)

	for rows.Next() {
		p := &model.Post{}

		if err := rows.Scan(&p.ID, &p.Author, &p.Title, &p.Text, &p.IsCommented, &p.CountComms, &p.LastCommentTime); err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		posts = append(posts, p)
	}

	for i, p := range posts {
		p, err = r.getAllUsersFromPost(p)

		if err != nil {
			config.ErrorLog(err)
			return nil, err
		}

		if p.IsCommented {
			p, err = r.getAllCommsFromPost(p)
			if err != nil {
				config.ErrorLog(err)
				return nil, err
			}
		}

		posts[i] = p
	}

	return posts, nil
}

func (r *PostRepo) InsertPost(post *model.Post) error {
	stmt := `INSERT INTO posts (author, title, text, isCommented, countComments, lastCommentTime)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(stmt, post.Author, post.Title, post.Text, post.IsCommented, post.CountComms, post.LastCommentTime)

	if err != nil {
		return err
	}

	config.InfoLog("new post inserted")
	return nil
}

func (r *PostRepo) GetPostById(id int64) (*model.Post, error) {
	stmt := `SELECT * FROM posts WHERE id = $1`

	row := r.db.QueryRow(stmt, id)

	post := &model.Post{}

	if err := row.Scan(&post.ID, &post.Author, &post.Title, &post.Text, &post.IsCommented, &post.CountComms, &post.LastCommentTime); err != nil {
		return nil, err
	}

	post, err := r.getAllUsersFromPost(post)

	if err != nil {
		return nil, err
	}

	post, err = r.getAllCommsFromPost(post)

	if err != nil {
		return nil, err
	}

	return post, nil

}

func (r *PostRepo) AddUserInPost(postId int64, userID int64) error {
	stmt := `INSERT INTO users_posts (user_id, post)
	VALUES ($1, $2)`

	_, err := r.db.Exec(stmt, userID, postId)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) UpdatePost(post *model.Post) error {
	stmt := `UPDATE posts
	SET title = $1, text = $2, isCommented = $3, countComments = $4, lastCommentTime = $5
	WHERE id = $6`

	_, err := r.db.Exec(stmt, post.Title, post.Text, post.IsCommented, post.CountComms, post.LastCommentTime, post.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) DeletePostById(id int64) error {
	stmt := `DELETE FROM posts WHERE id = $1`

	_, err := r.db.Exec(stmt, id)

	if err != nil {
		return err
	}

	return nil
}
