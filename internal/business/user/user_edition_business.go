package userbiz

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
	"log"
	"time"
)

type UserEditionStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, data *usersmodel.UserEdition) error
}

type editUserBusiness struct {
	store UserEditionStorage
}

func NewEditUserBiz(store UserEditionStorage) *editUserBusiness {
	return &editUserBusiness{
		store: store,
	}
}

func (biz *editUserBusiness) UpdateUser(ctx context.Context, id int, data *usersmodel.UserEdition) error {
	start := time.Now()
	_, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(usersmodel.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(usersmodel.EntityName, err)
	}

	if err := biz.store.UpdateUser(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(usersmodel.EntityName, err)
	}
	log.Printf("User update business took %v", time.Since(start))

	return nil
}
