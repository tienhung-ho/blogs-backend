package permissionbusiness

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
	"log"
)

type PermissionDeletionStorage interface {
	FindPermissions(ctx context.Context, cond map[string]interface{}) (*permissionmodel.Permission, error)
	DeletePermission(ctx context.Context, cond map[string]interface{}) error
}

type deletionPermissionBusiness struct {
	store PermissionDeletionStorage
}

func NewPermissionDeletionBiz(store PermissionDeletionStorage) *deletionPermissionBusiness {

	return &deletionPermissionBusiness{
		store: store,
	}
}

func (biz *deletionPermissionBusiness) DeletePermission(ctx context.Context, cond map[string]interface{}) error {

	record, err := biz.store.FindPermissions(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(permissionmodel.EntityName, err)
	}

	log.Printf("%v", record.Id)

	if err := biz.store.DeletePermission(ctx, map[string]interface{}{
		"id": record.Id,
	}); err != nil {
		return common.ErrCannotDeleteEntity(permissionmodel.EntityName, err)
	}

	return nil
}
