package permissionbusiness

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

type PermissionFinditionStorage interface {
	FindPermissions(ctx context.Context, cond map[string]interface{}) (*permissionmodel.Permission, error)
}

type permissionFinditionBusiness struct {
	store PermissionFinditionStorage
}

func NewPermissionFinditionBiz(store PermissionFinditionStorage) *permissionFinditionBusiness {
	return &permissionFinditionBusiness{
		store: store,
	}
}

func (biz *permissionFinditionBusiness) FindPermissions(ctx context.Context, cond map[string]interface{}) (*permissionmodel.Permission, error) {
	record, err := biz.store.FindPermissions(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(permissionmodel.EntityName, err)
	}

	return record, nil
}
