package blogcategorystorage

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

func (s *mysqlStorage) DeleteBlogCategory(ctx context.Context, cond map[string]interface{}) error {
	db := s.db.Begin()

	if err := db.Table(blogcategorymodel.BlogCategory{}.TableName()).
		Where(cond).Updates(map[string]interface{}{"deleted": true}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
