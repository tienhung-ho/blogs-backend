package permissionbusiness

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

type PermissionUpdationStorage interface {
	FindPermissions(ctx context.Context, cond map[string]interface{}) (*permissionmodel.Permission, error)
	UpdatePermission(ctx context.Context, cond map[string]interface{}, data *permissionmodel.PermissionUpdation) error
}

type permissionUpdationBusiness struct {
	store PermissionUpdationStorage
}

func NewPermissionUpdationBiz(store PermissionUpdationStorage) *permissionUpdationBusiness {
	return &permissionUpdationBusiness{
		store: store,
	}
}

func (biz *permissionUpdationBusiness) UpdatePermission(ctx context.Context, cond map[string]interface{}, data *permissionmodel.PermissionUpdation) error {
	record, err := biz.store.FindPermissions(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(permissionmodel.EntityName, err)
	}

	if err := biz.store.UpdatePermission(ctx, map[string]interface{}{"id": record.Id}, data); err != nil {
		return common.ErrCannotUpdateEntity(permissionmodel.EntityName, err)
	}
	return nil
}
