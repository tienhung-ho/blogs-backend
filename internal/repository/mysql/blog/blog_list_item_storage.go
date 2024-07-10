package blogstorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

func (s *mysqlStorage) ListItem(ctx context.Context, cond map[string]interface{}) ([]blogmodel.Blog, error) {
	var records []blogmodel.Blog
	db := s.db.Begin()

	if db.Error != nil {
		return nil, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Table(blogmodel.Blog{}.TableName()).
		Where(cond).Find(&records).Error; err != nil {

		db.Rollback()
		return nil, common.ErrDB(db.Error)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	return records, nil
}
