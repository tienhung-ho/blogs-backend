package blogcategorystorage

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

func (s *mysqlStorage) UpdateBlogCategory(ctx context.Context, cond map[string]interface{}, data blogcategorymodel.UpdateBlogCategory) error {

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
