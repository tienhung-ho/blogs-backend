package blogstorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

func (s *mysqlStorage) DeleteBlog(ctx context.Context, cond map[string]interface{}) error {
	db := s.db.Begin()

	if err := db.WithContext(ctx).Table(blogmodel.Blog{}.TableName()).Where(cond).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
