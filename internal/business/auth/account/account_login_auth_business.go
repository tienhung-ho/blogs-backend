package accountauthbiz

import (
	"blogs/internal/common"
	"blogs/internal/helpers"
	jwtcus "blogs/internal/helpers/token/jwt"
	accountmodel "blogs/internal/model/accounts"
	accountautmodel "blogs/internal/model/auth/account"
	"context"
	"errors"
)

type LoginOfAccountStorage interface {
	FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error)
}

type loginOfAccountBusiness struct {
	store      LoginOfAccountStorage
	jwtService *jwtcus.JwtServices
}

func NewLoginOfAccountBiz(store LoginOfAccountStorage, jwtService *jwtcus.JwtServices) *loginOfAccountBusiness {
	return &loginOfAccountBusiness{
		store:      store,
		jwtService: jwtService,
	}
}

func (biz *loginOfAccountBusiness) Login(ctx context.Context, login *accountautmodel.AccountLogin) (*accountmodel.SimpleAccount, error) {
	acc, err := biz.store.FindAccount(ctx, map[string]interface{}{"username": login.Username})

	if err != nil {
		return nil, common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}

	if acc.Deleted {
		return nil, common.ErrEntityDeleted(accountmodel.EntityName, err)
	}

	hasher := helpers.NewHashBcrypt(login.Password)

	if ok := hasher.ComparePass(acc.Password); !ok {
		return nil, common.ErrEmailOrPasswordInvalid(accountmodel.EntityName, errors.New("could not login"))
	}

	simpleAccount := acc.ToSimpleAccount()

	return simpleAccount, nil

}
