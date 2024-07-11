package rolestorage

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

func (s *mysqlStorage) FindRole(ctx context.Context, cond map[string]interface{}) (*rolemodel.Role, error) {

	var record rolemodel.Role

	db := s.db.Begin()

	if db.Error != nil {
		return nil, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.Table(rolemodel.Role{}.TableName()).WithContext(ctx).Preload("Permissions").Where(cond).Find(&record).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	return &record, nil
}
