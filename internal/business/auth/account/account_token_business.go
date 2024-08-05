package accountauthbiz

import (
	jwtcus "blogs/internal/helpers/token/jwt"
	accountmodel "blogs/internal/model/accounts"
	accountautmodel "blogs/internal/model/auth/account"
	"context"
	"log"
	"time"
)

type AccountTokenStorage interface {
	FindAccount(ctx context.Context, cond map[string]interface{}) (*accountmodel.Account, error)
}

type tokenBiz struct {
	store      AccountTokenStorage
	jwtService *jwtcus.JwtServices
}

func NewTokenBiz(store AccountTokenStorage, jwtService *jwtcus.JwtServices) *tokenBiz {
	return &tokenBiz{
		store:      store,
		jwtService: jwtService,
	}
}

type tokenResult struct {
	token string
	err   error
}

func (biz *tokenBiz) GenerateToken(ctx context.Context, accountId int, role string) (*accountautmodel.AccountToken, error) {
	account, err := biz.store.FindAccount(ctx, map[string]interface{}{"id": accountId})
	start := time.Now()

	if err != nil {
		return nil, err
	}

	// Channels to receive tokens and errors
	accessTokenChan := make(chan tokenResult)
	refreshTokenChan := make(chan tokenResult)

	// Generate refresh token
	go func() {
		token, err := biz.jwtService.GenerateTokenOfAccount(*account, 250*24*time.Hour, role)
		refreshTokenChan <- tokenResult{token: token, err: err}
	}()

	// Generate access token
	go func() {
		token, err := biz.jwtService.GenerateTokenOfAccount(*account, time.Hour, role)
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

	accountToken := accountautmodel.NewAccountToken(refreshToken, accessToken)
	log.Printf("Token business took %v", time.Since(start))
	return accountToken, nil
}
