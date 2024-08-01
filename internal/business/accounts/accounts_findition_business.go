package accountsbusiness

import (
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	"context"
)

type AccountFinditionStorage interface {
	FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error)
}

type accountFinditionBusiness struct {
	store AccountFinditionStorage
}

func NewAccountFinditionBiz(store AccountFinditionStorage) *accountFinditionBusiness {
	return &accountFinditionBusiness{
		store: store,
	}
}

func (biz *accountFinditionBusiness) FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error) {

	record, err := biz.store.FindAccount(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}

	return record, nil
}
