package rolestorage

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

func (s *mysqlStorage) CreateRoleWithPermissions(ctx context.Context, role *rolemodel.RoleCreation) (int, error) {
	db := s.db.Begin()

	if db.Error != nil {
		return 0, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Create(role).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	// Link role with permissions if any permissions are provided
	if len(role.Permissions) > 0 {
		if err := db.Model(role).Association("Permissions").Replace(role.Permissions); err != nil {
			db.Rollback()
			return 0, common.ErrDB(err)
		}
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return role.Id, nil
}
