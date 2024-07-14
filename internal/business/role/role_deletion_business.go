package rolebusiness

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

type RoleDeletionStorage interface {
	FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error)
	DeleteRole(ctx context.Context, cond map[string]interface{}) error
}

type roleDeletionBusiness struct {
	store RoleDeletionStorage
}

func NewRoleDeletionBusiness(store RoleDeletionStorage) *roleDeletionBusiness {
	return &roleDeletionBusiness{
		store: store,
	}
}

func (biz *roleDeletionBusiness) DeleteRole(ctx context.Context, cond map[string]interface{}) error {

	record, err := biz.store.FindRole(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(rolemodel.RoleEntityName, err)
	}

	if err := biz.store.DeleteRole(ctx, map[string]interface{}{"id": record.Id}); err != nil {
		return common.ErrCannotDeleteEntity(rolemodel.RoleEntityName, err)
	}

	return nil
}
