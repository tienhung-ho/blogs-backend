package rolebusiness

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
	"errors"
	"reflect"
)

type RoleUpdationStorage interface {
	FindPermissionsByName(ctx context.Context, names []string) ([]rolemodel.Permission, error)
	FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error)
	UpdateRole(ctx context.Context, cond map[string]interface{}, data rolemodel.RoleUpdation) error
}

type roleUpdationBusiness struct {
	store RoleUpdationStorage
}

func NewRoleUpdationBiz(store RoleUpdationStorage) *roleUpdationBusiness {
	return &roleUpdationBusiness{
		store: store,
	}
}

func (biz *roleUpdationBusiness) UpdateRole(ctx context.Context, cond map[string]interface{}, data rolemodel.RoleUpdation) error {
	record, err := biz.store.FindRole(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(rolemodel.RoleEntityName, err)
	}

	// Get permissions
	permissionNames := make([]string, len(data.Permissions))
	for i, perm := range data.Permissions {
		permissionNames[i] = perm.Name
	}

	permissions, err := biz.store.FindPermissionsByName(ctx, permissionNames)
	if err != nil {
		return err
	}

	// Update the role with the found permissions
	data.Permissions = permissions

	if reflect.DeepEqual(data, rolemodel.RoleUpdation{}) {
		return common.ErrInternal(errors.New("role data is empty"))
	}

	if err := biz.store.UpdateRole(ctx, map[string]interface{}{"name": record.Name}, data); err != nil {
		return common.ErrCannotUpdateEntity(rolemodel.RoleEntityName, err)
	}

	return nil
}
