package permissionstorage

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
)

func (s *mysqlStorage) CreatPermission(ctx context.Context, data *permissionmodel.PermissionCreation) (int, error) {

	db := s.db.Begin()

	if db.Error != nil {
		return 0, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return data.Id, nil

}
