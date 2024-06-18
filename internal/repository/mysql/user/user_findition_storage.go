package userstorage

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"

	"gorm.io/gorm"
)

func (s *mysqlStorage) GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error) {
	var data usersmodel.Users

	// Start timing the query
	// start := time.Now()

	// Perform the query
	if err := s.db.WithContext(ctx).Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	// Log the time taken to execute the query
	// log.Printf("GetUser query took %v", time.Since(start))

	return &data, nil
}
