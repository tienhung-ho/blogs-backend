package userbiz

import (
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	usersmodel "blogs/internal/model/users"
	"context"
)

type UserListItemStorage interface {
	ListItem(ctx context.Context, filter *filtermodel.Filter, paging *common.Paging, morekeys ...string) ([]usersmodel.Users, error)
}

type userListItemBusiness struct {
	store UserListItemStorage
}

func NewUserListItemBiz(store UserListItemStorage) *userListItemBusiness {
	return &userListItemBusiness{
		store: store,
	}
}

func (biz *userListItemBusiness) ListItem(ctx context.Context, filter *filtermodel.Filter, paging *common.Paging, morekeys ...string) ([]usersmodel.Users, error) {

	records, err := biz.store.ListItem(ctx, filter, paging, morekeys...)

	if err != nil {
		return nil, common.ErrCannotListEntity(usersmodel.EntityName, err)
	}

	return records, nil
}
