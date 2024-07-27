package permissionbusiness

import (
	permissionmodel "blogs/internal/model/permission"
	"context"
)

type PermissionListItemStorage interface {
	ListPermissionsByName(ctx context.Context, cond map[string]interface{}) ([]permissionmodel.Permission, error)
}
