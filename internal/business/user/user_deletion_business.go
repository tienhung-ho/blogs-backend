package userbiz

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

type DeleteUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
	DeleteUser(ctx context.Context, cond map[string]interface{}) error
}

type deleteUserBusiness struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteUserBusiness {
	return &deleteUserBusiness{
		store: store,
	}
}

func (biz *deleteUserBusiness) DeleteUser(ctx context.Context, id int) error {

	user, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(usersmodel.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(usersmodel.EntityName, err)
	}

	if user.Deleted {
		return common.ErrEntityDeleted(usersmodel.EntityName, err)
	}

	if err := biz.store.DeleteUser(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(usersmodel.EntityName, err)
	}
	return nil
}
