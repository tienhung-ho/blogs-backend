package accountsstogare

import (
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	"context"
)

func (s *mysqlStorage) FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error) {

	var record accountmodel.Account

	db := s.db
	if err := db.WithContext(ctx).Where(cond).First(&record).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &record, nil
}
