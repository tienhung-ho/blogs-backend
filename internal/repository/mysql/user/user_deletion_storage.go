package userstorage

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

func (s *mysqlStorage) DeleteUser(ctx context.Context, cond map[string]interface{}) error {
	db := s.db.Begin()

	if err := db.Table(usersmodel.Users{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"deleted": true,
	}).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
