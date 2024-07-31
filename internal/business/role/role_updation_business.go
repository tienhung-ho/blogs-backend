package rolebusiness

import (
	permissionbusiness "blogs/internal/business/permission"
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
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
	roleStore       RoleUpdationStorage
	permissionStore permissionbusiness.PermissionListItemStorage
}

func NewRoleUpdationBiz(roleStore RoleUpdationStorage, permissionStore permissionbusiness.PermissionListItemStorage) *roleUpdationBusiness {
	return &roleUpdationBusiness{
		roleStore:       roleStore,
		permissionStore: permissionStore,
	}
}

func (biz *roleUpdationBusiness) UpdateRole(ctx context.Context, cond map[string]interface{}, data rolemodel.RoleUpdation) error {
	record, err := biz.roleStore.FindRole(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(rolemodel.RoleEntityName, err)
	}

	// Get permissions
	permissionNames := make([]string, len(data.Permissions))
	for i, perm := range data.Permissions {
		permissionNames[i] = perm.Name
	}

	permissionCond := map[string]interface{}{
		"names": permissionNames,
	}

	permissions, err := biz.permissionStore.ListPermissions(ctx, permissionCond, &common.Paging{}, &filtermodel.Filter{})
	if err != nil {
		return err
	}

	// Update the role with the found permissions
	data.Permissions = permissions

	if reflect.DeepEqual(data, rolemodel.RoleUpdation{}) {
		return common.ErrInternal(errors.New("role data is empty"))
	}

	if err := biz.roleStore.UpdateRole(ctx, map[string]interface{}{"name": record.Name}, data); err != nil {
		return common.ErrCannotUpdateEntity(rolemodel.RoleEntityName, err)
	}

	return nil
}
