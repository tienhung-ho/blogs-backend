package blogstorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

func (s *mysqlStorage) UpdateBlog(ctx context.Context, cond map[string]interface{}, data *blogmodel.BlogUpdate, morekeys ...string) error {

	db := s.db.Begin()

	if err := db.WithContext(ctx).Where(cond).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
