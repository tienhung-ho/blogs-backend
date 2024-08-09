package permissionstorage

import (
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	permissionmodel "blogs/internal/model/permission"
	"context"

	"gorm.io/gorm"
)

func (s *mysqlStorage) ListPermissions(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekeys ...string) ([]permissionmodel.Permission, error) {
	var permissions []permissionmodel.Permission

	db := s.db.WithContext(ctx)

	// // Đếm tổng số lượng items
	if err := s.countPermissions(db, cond, filter, paging); err != nil {
		return nil, err
	}

	// Xây dựng truy vấn động
	query := s.buildQuery(db, cond, filter)

	// Thêm phân trang
	query = s.addPaging(query, paging)

	// Thực hiện truy vấn
	if err := query.Find(&permissions).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return permissions, nil
}

func (s *mysqlStorage) countPermissions(db *gorm.DB, cond map[string]interface{}, filter *filtermodel.Filter, paging *common.Paging) error {
	if names, ok := cond["names"]; ok {
		db = db.Where("name IN ?", names)
	} else {
		db = db.Where(cond)
		if filter != nil && filter.Status != "" {
			db = db.Where("status = ?", filter.Status)
		}
	}
	if err := db.Table(permissionmodel.Permission{}.TableName()).Count(&paging.Total).Error; err != nil {
		return common.NewErrorResponse(err, "Error count items from database", err.Error(), "CouldNotCount")
	}
	return nil
}

func (s *mysqlStorage) buildQuery(db *gorm.DB, cond map[string]interface{}, filter *filtermodel.Filter) *gorm.DB {
	if names, ok := cond["names"]; ok {
		db = db.Where("name IN ?", names)
	} else {
		db = db.Where(cond)
		if filter != nil && filter.Status != "" {
			db = db.Where("status = ?", filter.Status)
		}
	}
	return db
}

func (s *mysqlStorage) addPaging(db *gorm.DB, paging *common.Paging) *gorm.DB {
	return db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit)
}
