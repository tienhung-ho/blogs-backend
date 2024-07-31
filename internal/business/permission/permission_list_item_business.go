package permissionbusiness

import (
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

type PermissionListItemStorage interface {
	ListPermissions(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekeys ...string) ([]permissionmodel.Permission, error)
}

type permissionListItemBusiness struct {
	store PermissionListItemStorage
}

func NewPermissionListItem(store PermissionListItemStorage) *permissionListItemBusiness {
	return &permissionListItemBusiness{
		store: store,
	}
}

func (biz *permissionListItemBusiness) ListPermissions(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekeys ...string) ([]permissionmodel.Permission, error) {
	records, err := biz.store.ListPermissions(ctx, cond, paging, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(permissionmodel.EntityName, err)
	}

	return records, nil
}
