package rolebusiness

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

type RoleListItemStorage interface {
	ListItem(ctx context.Context, cond map[string]interface{}) ([]rolemodel.Role, error)
}

type roleListItemBusiness struct {
	store RoleListItemStorage
}

func NewRoleListItemBiz(store RoleListItemStorage) *roleListItemBusiness {
	return &roleListItemBusiness{
		store: store,
	}
}

func (biz *roleListItemBusiness) ListItem(ctx context.Context, cond map[string]interface{}) ([]rolemodel.Role, error) {

	records, err := biz.store.ListItem(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotListEntity(rolemodel.RoleEntityName, err)
	}

	return records, nil
}
