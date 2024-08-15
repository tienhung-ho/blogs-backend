package blogstorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	filtermodel "blogs/internal/model/filter"
	"context"
	"gorm.io/gorm"
)

func (s *mysqlStorage) ListItem(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekeys ...string) ([]blogmodel.Blog, error) {
	var records []blogmodel.Blog
	db := s.db.Begin()

	if db.Error != nil {
		return nil, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	// // Đếm tổng số lượng items
	if err := s.countBlog(db, cond, filter, paging); err != nil {
		return nil, err
	}

	// Xây dựng truy vấn động
	query := s.buildQuery(db, cond, filter)

	// Thêm phân trang
	query = s.addPaging(query, paging)

	// Thực hiện truy vấn
	if err := query.Find(&records).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := query.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	return records, nil
}

func (s *mysqlStorage) countBlog(db *gorm.DB, cond map[string]interface{}, filter *filtermodel.Filter, paging *common.Paging) error {
	if names, ok := cond["names"]; ok {
		db = db.Where("name IN ?", names)
	} else {
		db = db.Where(cond)
		if filter != nil && filter.Status != "" {
			db = db.Where("status = ?", filter.Status)
		}
	}
	if err := db.Table(blogmodel.Blog{}.TableName()).Count(&paging.Total).Error; err != nil {
		return common.NewErrorResponse(err, "Error count items from database", err.Error(), "CouldNotCount")
	}
	return nil
}

func (s *mysqlStorage) buildQuery(db *gorm.DB, cond map[string]interface{}, filter *filtermodel.Filter) *gorm.DB {
	db = db.Where(cond)
	if filter != nil && filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}
	return db
}

func (s *mysqlStorage) addPaging(db *gorm.DB, paging *common.Paging) *gorm.DB {
	return db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit)
}
