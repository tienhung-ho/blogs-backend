package rolestorage

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

func (s *mysqlStorage) DeleteRole(ctx context.Context, cond map[string]interface{}) error {

	db := s.db.Begin()

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.Table(rolemodel.Role{}.TableName()).Where(cond).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(db.Error)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(db.Error)
	}

	return nil
}
