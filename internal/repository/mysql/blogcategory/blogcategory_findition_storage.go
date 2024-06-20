package blogcategorystorage

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"

	"gorm.io/gorm"
)

func (s *mysqlStorage) GetBlogCategory(ctx context.Context, cond map[string]interface{}) (*blogcategorymodel.BlogCategory, error) {
	var data blogcategorymodel.BlogCategory
	if err := s.db.WithContext(ctx).Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
