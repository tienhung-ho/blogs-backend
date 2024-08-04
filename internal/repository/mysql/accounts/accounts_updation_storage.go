package accountsstogare

import (
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	"context"
)

func (s *mysqlStorage) UpdateAccount(ctx context.Context, cond map[string]interface{}, data *accountmodel.AccountUpdation) error {

	db := s.db.Begin()

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Where(cond).Updates(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
