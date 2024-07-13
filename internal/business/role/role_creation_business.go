package rolebusiness

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
	"errors"
)

type RoleCreationStorage interface {
	FindPermissionsByName(ctx context.Context, names []string) ([]rolemodel.Permission, error)
	FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error)
	CreateRoleWithPermissions(ctx context.Context, role *rolemodel.RoleCreation) (int, error)
}

type roleCreationBusiness struct {
	store RoleCreationStorage
}

func NewRoleCreationBiz(store RoleCreationStorage) *roleCreationBusiness {
	return &roleCreationBusiness{
		store: store,
	}
}

func (biz *roleCreationBusiness) CreateRole(ctx context.Context, data rolemodel.RoleCreation) (int, error) {
	record, err := biz.store.FindRole(ctx, map[string]interface{}{"name": data.Name})

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
	permissions, err := biz.store.FindPermissionsByName(ctx, permissionNames)
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

	roleId, err := biz.store.CreateRoleWithPermissions(ctx, &role)
	if err != nil {
		return 0, err
	}

	return roleId, nil
}
