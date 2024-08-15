package blogstorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

func (s *mysqlStorage) CreateBlog(ctx context.Context, data *blogmodel.BlogCreation, morekeys ...string) (int, error) {

	db := s.db.Begin()

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return data.Id, nil
}
