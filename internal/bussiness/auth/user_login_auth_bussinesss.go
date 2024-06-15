package authbiz

import (
	"blogs/internal/common"
	"blogs/internal/helpers"
	jwtcus "blogs/internal/helpers/token/jwt"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	"context"
	"errors"
	"time"
)

type UserLoginStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
}

type loginUserBussiness struct {
	store      UserLoginStorage
	jwtService *jwtcus.JwtServices
}

func NewLoginUserBiz(user UserLoginStorage, jwtService *jwtcus.JwtServices) *loginUserBussiness {
	return &loginUserBussiness{
		store:      user,
		jwtService: jwtService,
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

	accessToken, err := biz.jwtService.GenerateToken(*user, time.Hour)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.jwtService.GenerateToken(*user, 250*24*time.Hour)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := authmodel.NewUserToken(accessToken, refreshToken)

	return account, nil
}