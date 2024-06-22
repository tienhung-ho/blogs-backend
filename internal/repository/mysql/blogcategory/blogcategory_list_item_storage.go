package blogcategorystorage

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

func (s *mysqlStorage) ListItem(ctx context.Context, cond map[string]interface{}) ([]blogcategorymodel.BlogCategory, error) {

	var data []blogcategorymodel.BlogCategory

	err := s.db.Table(blogcategorymodel.BlogCategory{}.TableName()).WithContext(ctx).Where(cond).Find(&data).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
