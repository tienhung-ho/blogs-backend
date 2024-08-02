package accountsbusiness

import (
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	"context"
)

type AccountDeletionStorage interface {
	DeleteAccount(ctx context.Context, cond map[string]interface{}) error
	FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error)
}

type accountDeletionBusiness struct {
	store AccountDeletionStorage
}

func NewAccountDeletionBiz(store AccountDeletionStorage) *accountDeletionBusiness {
	return &accountDeletionBusiness{
		store: store,
	}
}

func (biz *accountDeletionBusiness) DeleteAccount(ctx context.Context, cond map[string]interface{}) error {
	record, err := biz.store.FindAccount(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}

	if err := biz.store.DeleteAccount(ctx, map[string]interface{}{"id": record.ID}); err != nil {
		return common.ErrCannotDeleteEntity(accountmodel.EntityName, err)
	}

	return nil
}
