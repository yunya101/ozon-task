package data

import "github.com/yunya101/ozon-task/internal/model"

type UserRepoInMem struct {
	users   map[int64]*model.User
	countId int64
}

func NewUserRepoInMem() *UserRepoInMem {
	return &UserRepoInMem{
		users:   map[int64]*model.User{},
		countId: 1,
	}
}

func (r *UserRepoInMem) Insert(user *model.User) error {
	user.ID = r.countId
	r.users[r.countId] = user

	r.countId++

	return nil
}

func (r *UserRepoInMem) DeleteById(id int64) error {

	delete(r.users, id)

	return nil
}
