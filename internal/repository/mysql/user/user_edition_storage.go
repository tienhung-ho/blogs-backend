package userstorage

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

func (sql *mysqlStorage) UpdateUser(ctx context.Context, cond map[string]interface{}, data *usersmodel.UserEdition) error {
	db := sql.db.Begin()

	if err := db.Where(cond).Updates(data).Error; err != nil {
		db.Rollback()
		return common.NewErrorResponse(err, "Error update items from database", err.Error(), "ErrorFetching")
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
