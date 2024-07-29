package permissionstorage

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

func (s *mysqlStorage) UpdatePermission(ctx context.Context, cond map[string]interface{}, data *permissionmodel.PermissionUpdation) error {

	db := s.db.Begin()

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Where(cond).Updates(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(db.Error)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(db.Error)
	}

	return nil
}
