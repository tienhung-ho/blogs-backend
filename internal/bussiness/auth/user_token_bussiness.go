package authbiz

import (
	jwtcus "blogs/internal/helpers/token/jwt"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	"context"
	"log"
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

type tokenResult struct {
	token string
	err   error
}

func (biz *tokenBiz) GenerateToken(ctx context.Context, userId int) (*authmodel.UserToken, error) {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"id": userId})
	start := time.Now()

	if err != nil {
		return nil, err
	}

	// Channels to receive tokens and errors
	accessTokenChan := make(chan tokenResult)
	refreshTokenChan := make(chan tokenResult)

	// Generate refresh token
	go func() {
		token, err := biz.jwtService.GenerateToken(*user, 250*24*time.Hour)
		refreshTokenChan <- tokenResult{token: token, err: err}
	}()

	// Generate access token
	go func() {
		token, err := biz.jwtService.GenerateToken(*user, time.Hour)
		accessTokenChan <- tokenResult{token: token, err: err}
	}()

	// Collect results
	var accessToken, refreshToken string
	for i := 0; i < 2; i++ {
		select {
		case res := <-accessTokenChan:
			if res.err != nil {
				return nil, res.err
			}
			accessToken = res.token
		case res := <-refreshTokenChan:
			if res.err != nil {
				return nil, res.err
			}
			refreshToken = res.token
		}
	}

	userToken := authmodel.NewUserToken(accessToken, refreshToken)
	log.Printf("Token business took %v", time.Since(start))
	return userToken, nil
}
