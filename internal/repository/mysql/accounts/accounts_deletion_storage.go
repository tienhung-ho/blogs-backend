package accountsstogare

import (
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	"context"
)

func (s *mysqlStorage) DeleteAccount(ctx context.Context, cond map[string]interface{}) error {

	db := s.db.Begin()

	if db.Error != nil {
		return common.ErrDB(db.Error)
	}

	defer common.RecoverTransaction(db)

	if err := db.WithContext(ctx).Table(accountmodel.Account{}.TableName()).
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
