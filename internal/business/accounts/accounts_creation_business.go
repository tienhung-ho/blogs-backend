package accountsbusiness

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	rolemodel "blogs/internal/model/role"
	"context"
	"reflect"
)

type AccountCreationStorage interface {
	CreateAccount(ctx context.Context, data *accountmodel.Account) error
}

type accountCreationBusiness struct {
	accountCreationStore AccountCreationStorage
	roleFinditionStore   rolebusiness.RoleFiditionStorage
}

func NewAccountCreationBiz(accountCreationStore AccountCreationStorage, roleFinditionStore rolebusiness.RoleFiditionStorage) *accountCreationBusiness {
	return &accountCreationBusiness{
		accountCreationStore: accountCreationStore,
		roleFinditionStore:   roleFinditionStore,
	}
}

func (biz *accountCreationBusiness) CreateAccount(ctx context.Context, data *accountmodel.Account) error {

	role, err := biz.roleFinditionStore.FindRole(ctx, map[string]interface{}{
		"deleted": false,
		"id":      data.RoleID,
	})

	if err != nil {
		return common.ErrNotFoundEntity(rolemodel.RoleEntityName, err)
	}

	if reflect.DeepEqual(role, rolemodel.Role{}) {
		return common.ErrNotFoundEntity(rolemodel.RoleEntityName, err)
	}

	if err := biz.accountCreationStore.CreateAccount(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}

	return nil
}
