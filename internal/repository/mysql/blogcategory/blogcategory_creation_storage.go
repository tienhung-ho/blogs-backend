package blogcategorystorage

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

func (s *mysqlStorage) CreateBlogCategory(ctx context.Context, data *blogcategorymodel.CreationBlogCategory) (int, error) {
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
