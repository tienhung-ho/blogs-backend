package authbiz

import (
	"blogs/internal/common"
	"blogs/internal/helpers"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	"context"
	"errors"
)

type UserLogin interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
}

type loginUserBussiness struct {
	store UserLogin
}

func NewLoginUserBiz(user UserLogin) *loginUserBussiness {
	return &loginUserBussiness{
		store: user,
	}
}

func (biz *loginUserBussiness) Login(ctx context.Context, loginUser *authmodel.UserLogin) (*authmodel.UserToken, error) {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"email": loginUser.Email})

	if err != nil {
		return nil, err
	}

	hasher := helpers.NewHashBcrypt(loginUser.Password)

	if ok := hasher.ComparePass(user.Password); !ok {
		return nil, common.ErrEmailOrPasswordInvalid(usersmodel.EntityName, errors.New("could not login"))
	}

	return nil, nil
}
