package userbiz

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

type FindUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
}

type findUserBusiness struct {
	db FindUserStorage
}

func NewFindUserBiz(db FindUserStorage) *findUserBusiness {
	return &findUserBusiness{db: db}
}

func (biz *findUserBusiness) GetItemById(ctx context.Context, id int) (*usersmodel.Users, error) {
	data, err := biz.db.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(usersmodel.EntityName, err)
	}

	return data, nil
}

func (biz *findUserBusiness) GetItemByCondition(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error) {
	data, err := biz.db.GetUser(ctx, cond)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usersmodel.EntityName, err)
	}

	return data, nil
}
