package data

import "github.com/yunya101/ozon-task/internal/model"

type CommRepoInMem struct {
	comms   map[int64]*model.Comment
	countId int64
}

func NewCommRepoInMem() *CommRepoInMem {
	return &CommRepoInMem{
		comms:   map[int64]*model.Comment{},
		countId: 1,
	}
}

func (r *CommRepoInMem) Insert(com *model.Comment) (int64, error) {
	com.ID = r.countId

	r.comms[r.countId] = com

	if com.ParentID > 0 {
		r.comms[com.ParentID].Comments = append(r.comms[com.ParentID].Comments, com)
	}

	r.countId++

	return com.ID, nil
}
