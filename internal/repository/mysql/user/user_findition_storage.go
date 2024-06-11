package userstorage

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"

	"gorm.io/gorm"
)

func (s *mysqlStorage) GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error) {
	var data usersmodel.Users

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
