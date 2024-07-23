package rolestorage

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

func (s *mysqlStorage) UpdateRole(ctx context.Context, cond map[string]interface{}, data rolemodel.RoleUpdation) error {

	db := s.db.Begin()

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Model(&rolemodel.Role{}).Where(cond).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	var role rolemodel.Role

	if err := db.WithContext(ctx).Where("name = ?", data.Name).First(&role).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Model(&role).Association("Permissions").Replace(data.Permissions); err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
