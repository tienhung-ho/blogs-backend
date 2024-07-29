package permissionstorage

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	"context"
	"log"
)

func (s *mysqlStorage) DeletePermission(ctx context.Context, cond map[string]interface{}) error {
	db := s.db.Begin()

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Table(permissionmodel.Permission{}.TableName()).
		Where(cond).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	var record permissionmodel.Permission

	if err := db.WithContext(ctx).Table(permissionmodel.Permission{}.TableName()).Where(cond).First(&record).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.WithContext(ctx).Model(&record).Association("Roles").Clear(); err != nil {
		log.Printf("%v", err)
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
