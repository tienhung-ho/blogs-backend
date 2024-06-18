package authbiz

import (
	"blogs/internal/common"
	"blogs/internal/helpers"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	"context"
)

type CreateUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*usersmodel.Users, error)
	CreateUser(ctx context.Context, data *usersmodel.UserCreation) (int, error)
}

type createUserBusiness struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBusiness {
	return &createUserBusiness{store: store}
}

func (biz *createUserBusiness) CreateUser(ctx context.Context, data *authmodel.UserRegister) (int, error) {
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return 0, common.ErrRecordExist(usersmodel.EntityName, err)
	}

	userName := data.Username

	if userName == "" {
		common.NewErrorResponse(usersmodel.ErrUserNameBlank, "Can not leave the user name blank", "User name empty!", "UserNameEmpty")
		return 0, usersmodel.ErrUserNameBlank
	}

	hasher := helpers.NewHashBcrypt(data.Password)

	hashedPassword, err := hasher.GeneratePass()

	if err != nil {
		return 0, err
	}

	userId, err := biz.store.CreateUser(ctx, data.DoRegister(hashedPassword))

	if err != nil {
		return 0, common.ErrCannotCreateEntity(usersmodel.EntityName, err)
	}

	return userId, nil
}
