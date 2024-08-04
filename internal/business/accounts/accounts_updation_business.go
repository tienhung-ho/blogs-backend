package accountsbusiness

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	rolemodel "blogs/internal/model/role"
	"context"
	"errors"
	"reflect"
)

type AccountUpdationStorage interface {
	UpdateAccount(ctx context.Context, cond map[string]interface{}, data *accountmodel.AccountUpdation) error
	FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error)
}

type accountUpdationBusiness struct {
	accountUpdationStorage AccountUpdationStorage
	roleFinditionStore     rolebusiness.RoleFiditionStorage
}

func NewAccountUpdationBiz(accountUpdationStorage AccountUpdationStorage, roleFinditionStore rolebusiness.RoleFiditionStorage) *accountUpdationBusiness {
	return &accountUpdationBusiness{
		accountUpdationStorage: accountUpdationStorage,
		roleFinditionStore:     roleFinditionStore,
	}
}

func (biz *accountUpdationBusiness) UpdateAccount(ctx context.Context, cond map[string]interface{}, data *accountmodel.AccountUpdation) error {

	record, err := biz.accountUpdationStorage.FindAccount(ctx, cond)

	if err != nil {
		return common.RecordNotFound
	}

	exists, isZero := common.CheckFieldExistsAndNotZero(data, "RoleID")

	if exists {
		if isZero {
			return common.ErrCannotCreateEntity(accountmodel.EntityName, errors.New("number is zero value"))
		} else {

			role, err := biz.roleFinditionStore.FindRole(ctx, map[string]interface{}{"id": data.RoleID})

			if err != nil {
				return common.RecordNotFound
			}

			if role.Id == 0 {
				return common.ErrCannotGetEntity(rolemodel.RoleEntityName, errors.New("role is empty"))
			}

			if reflect.DeepEqual(role, rolemodel.Role{}) {
				return common.ErrCannotGetEntity(rolemodel.RoleEntityName, errors.New("role is empty"))
			}
		}
	}

	if err := biz.accountUpdationStorage.UpdateAccount(ctx, map[string]interface{}{"username": record.Username}, data); err != nil {
		return common.ErrCannotUpdateEntity(accountmodel.EntityName, err)
	}

	return nil
}
