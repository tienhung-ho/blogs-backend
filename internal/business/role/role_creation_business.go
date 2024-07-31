package rolebusiness

import (
	permissionbusiness "blogs/internal/business/permission"
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	rolemodel "blogs/internal/model/role"
	"context"
	"errors"
)

type RoleCreationStorage interface {
	FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error)
	CreateRoleWithPermissions(ctx context.Context, role *rolemodel.RoleCreation) (int, error)
}

type roleCreationBusiness struct {
	roleStore       RoleCreationStorage
	permissionStore permissionbusiness.PermissionListItemStorage
}

func NewRoleCreationBiz(roleStore RoleCreationStorage, permissionStore permissionbusiness.PermissionListItemStorage) *roleCreationBusiness {
	return &roleCreationBusiness{
		roleStore:       roleStore,
		permissionStore: permissionStore,
	}
}

func (biz *roleCreationBusiness) CreateRole(ctx context.Context, data rolemodel.RoleCreation) (int, error) {
	record, err := biz.roleStore.FindRole(ctx, map[string]interface{}{"name": data.Name})

	if err != nil {
		return 0, common.ErrCannotGetEntity(rolemodel.RoleEntityName, err)
	}

	if record.Name != "" {
		return 0, common.ErrRecordExist(rolemodel.RoleEntityName, errors.New("record exist"))
	}

	// Lấy tất cả các tên permission từ dữ liệu đầu vào
	permissionNames := make([]string, len(data.Permissions))
	for i, perm := range data.Permissions {
		permissionNames[i] = perm.Name
	}

	// Tìm tất cả các permissions tồn tại trong database
	cond := map[string]interface{}{
		"names": permissionNames,
	}

	permissions, err := biz.permissionStore.ListPermissions(ctx, cond, &common.Paging{}, &filtermodel.Filter{})
	if err != nil {
		return 0, err
	}

	// Kiểm tra xem có thiếu permission nào không
	if len(permissions) != len(permissionNames) {
		return 0, common.ErrInvalidRequest(errors.New("one or more permissions do not exist"))
	}

	role := rolemodel.RoleCreation{
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
		Permissions: permissions,
	}

	roleId, err := biz.roleStore.CreateRoleWithPermissions(ctx, &role)
	if err != nil {
		return 0, err
	}

	return roleId, nil
}
