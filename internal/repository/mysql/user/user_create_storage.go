package userstorage

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

func (s *mysqlStorage) CreateUser(ctx context.Context, data *usersmodel.UserCreation) (int, error) {
	db := s.db.Begin()
	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}
	return data.Id, nil
}
