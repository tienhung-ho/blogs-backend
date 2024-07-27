package permissionstorage

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

// FindPermissions finds permissions based on various conditions
func (s *mysqlStorage) FindPermissions(ctx context.Context, cond map[string]interface{}) (*permissionmodel.Permission, error) {
	var permissions permissionmodel.Permission

	db := s.db

	if err := db.WithContext(ctx).Where(cond).First(&permissions).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &permissions, nil
}
