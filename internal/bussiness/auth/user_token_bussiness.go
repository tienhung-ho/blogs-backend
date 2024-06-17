package authbiz

import (
	jwtcus "blogs/internal/helpers/token/jwt"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	"context"
	"time"
)

type UserTokenStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
}

type tokenBiz struct {
	store      UserTokenStorage
	jwtService *jwtcus.JwtServices
}

func NewTokenBiz(store UserTokenStorage, jwtService *jwtcus.JwtServices) *tokenBiz {
	return &tokenBiz{
		store:      store,
		jwtService: jwtService,
	}
}

func (biz *tokenBiz) GenerateToken(ctx context.Context, userId int) (*authmodel.UserToken, error) {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"id": userId})

	if err != nil {
		return nil, err
	}

	refreshToken, err := biz.jwtService.GenerateToken(*user, 250*24*time.Hour)
	if err != nil {
		return nil, err
	}
	accessToken, err := biz.jwtService.GenerateToken(*user, time.Hour)
	if err != nil {
		return nil, err
	}

	userToken := authmodel.NewUserToken(accessToken, refreshToken)

	return userToken, nil
}
