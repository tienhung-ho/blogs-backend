package rolestorage

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

func (s *mysqlStorage) ListItem(ctx context.Context, cond map[string]interface{}) ([]rolemodel.Role, error) {

	var records []rolemodel.Role
	db := s.db.Begin()

	if db.Error != nil {
		return nil, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Where(cond).Preload("Permissions").Find(&records).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(db.Error)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(db.Error)
	}

	return records, nil

}
