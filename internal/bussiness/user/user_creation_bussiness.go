package userbiz

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"context"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *usersmodel.UserCreation) error
}

type createUserBussiness struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBussiness {
	return &createUserBussiness{store: store}
}

func (biz *createUserBussiness) CreateUser(ctx context.Context, data *usersmodel.UserCreation) error {

	userName := data.Username

	if userName == "" {
		common.NewErrorResponse(usersmodel.ErrUserNameBlank, "Can not leave the user name blank", "User name empty!", "UserNameEmpty")
		return usersmodel.ErrUserNameBlank
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usersmodel.EntityName, err)
	}

	return nil
}
