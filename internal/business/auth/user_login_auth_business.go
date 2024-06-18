package authbiz

import (
	"blogs/internal/common"
	"blogs/internal/helpers"
	jwtcus "blogs/internal/helpers/token/jwt"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	"context"
	"errors"
	"log"
	"time"
)

type UserLoginStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
}

type loginUserBusiness struct {
	store      UserLoginStorage
	jwtService *jwtcus.JwtServices
}

func NewLoginUserBiz(user UserLoginStorage, jwtService *jwtcus.JwtServices) *loginUserBusiness {
	return &loginUserBusiness{
		store:      user,
		jwtService: jwtService,
	}
}

func (biz *loginUserBusiness) Login(ctx context.Context, loginUser *authmodel.UserLogin) (*usersmodel.SimpleUser, error) {
	start := time.Now()
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"email": loginUser.Email})

	if err != nil {
		return nil, err
	}

	hasher := helpers.NewHashBcrypt(loginUser.Password)

	if ok := hasher.ComparePass(user.Password); !ok {
		return nil, common.ErrEmailOrPasswordInvalid(usersmodel.EntityName, errors.New("could not login"))
	}

	simpleUser := user.ToSimpleUser()

	log.Printf("User business took %v", time.Since(start))

	return simpleUser, nil
}
