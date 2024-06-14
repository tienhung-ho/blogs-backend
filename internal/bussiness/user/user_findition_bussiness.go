package userbiz

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

type FindUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
}

type findUserBussiness struct {
	db FindUserStorage
}

func NewFindUserBiz(db FindUserStorage) *findUserBussiness {
	return &findUserBussiness{db: db}
}

func (biz *findUserBussiness) GetItemById(ctx context.Context, id int) (*usersmodel.Users, error) {
	data, err := biz.db.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(usersmodel.EntityName, err)
	}

	return data, nil
}

func (biz *findUserBussiness) GetItemByCondition(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error) {
	data, err := biz.db.GetUser(ctx, cond)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usersmodel.EntityName, err)
	}

	return data, nil
}
