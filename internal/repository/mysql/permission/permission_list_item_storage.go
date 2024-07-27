package permissionstorage

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

// FindPermissions finds permissions based on various conditions
func (s *mysqlStorage) ListPermissionsByName(ctx context.Context, cond map[string]interface{}) ([]permissionmodel.Permission, error) {
	var permissions []permissionmodel.Permission

	// Sử dụng transaction nếu cần thiết
	db := s.db.WithContext(ctx)

	// Xây dựng truy vấn động
	if names, ok := cond["names"]; ok {
		db = db.Where("name IN ?", names)
	}

	if err := db.Find(&permissions).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return permissions, nil
}
