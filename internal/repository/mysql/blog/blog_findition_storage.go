package blogstorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"

	"gorm.io/gorm"
)

func (s *mysqlStorage) GetBlog(ctx context.Context, cond map[string]interface{}, morekeys ...string) (*blogmodel.Blog, error) {

	db := s.db.Begin()

	var record blogmodel.Blog

	if err := db.WithContext(ctx).Where(cond).First(&record).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			db.Rollback()
			return nil, common.RecordNotFound
		}

		db.Rollback()
		return nil, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	return &record, nil
}
