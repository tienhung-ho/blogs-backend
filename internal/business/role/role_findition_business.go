package rolebusiness

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

type RoleFiditionStorage interface {
	FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error)
}

type roleFiditionBusiness struct {
	store RoleFiditionStorage
}

func NewRoleFiditionBiz(store RoleFiditionStorage) *roleFiditionBusiness {
	return &roleFiditionBusiness{
		store: store,
	}
}

func (biz *roleFiditionBusiness) FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error) {

	role, err := biz.store.FindRole(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(rolemodel.RoleEntityName, err)
	}

	return role, nil
}
