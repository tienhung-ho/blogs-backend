package rolestorage

import (
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	"context"
)

func (s *mysqlStorage) FindPermissionsByName(ctx context.Context, names []string) ([]rolemodel.Permission, error) {
	var permissions []rolemodel.Permission
	if err := s.db.WithContext(ctx).Where("name IN ?", names).Find(&permissions).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return permissions, nil
}
