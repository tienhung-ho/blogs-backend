package userstorage

import (
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	usersmodel "blogs/internal/model/users"
	"context"
)

func (s *mysqlStorage) ListItem(ctx context.Context, filter *filtermodel.Filter, paging *common.Paging, morekeys ...string) ([]usersmodel.Users, error) {

	var records []usersmodel.Users
	db := s.db.Begin()

	if db.Error != nil {
		return nil, common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", f.Status)
		}
	}

	if err := db.Table(usersmodel.Users{}.TableName()).Count(&paging.Total).Error; err != nil {
		db.Rollback()
		return nil, common.NewErrorResponse(err, "Error count items from database", err.Error(), "CouldNotCount")
	}

	if err := db.WithContext(ctx).Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&records).Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, common.ErrDB(err)
	}

	return records, nil
}
