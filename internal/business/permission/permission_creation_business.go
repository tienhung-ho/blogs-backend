package permissionbusiness

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
	"errors"
	"reflect"
)

type PermissionCreationStorage interface {
	CreatPermission(ctx context.Context, data *permissionmodel.PermissionCreation) (int, error)
}

type permissionCreationBusiness struct {
	store PermissionCreationStorage
}

func NewPermissionCreationBiz(store PermissionCreationStorage) *permissionCreationBusiness {
	return &permissionCreationBusiness{
		store: store,
	}
}

func (biz *permissionCreationBusiness) CreatPermission(ctx context.Context, data *permissionmodel.PermissionCreation) (int, error) {

	if reflect.DeepEqual(*data, permissionmodel.PermissionCreation{}) {
		return 0, common.ErrInternal(errors.New("role data is empty"))
	}

	recordId, err := biz.store.CreatPermission(ctx, data)

	if err != nil {
		return 0, common.ErrCannotCreateEntity(permissionmodel.EntityName, err)
	}

	return recordId, nil
}
