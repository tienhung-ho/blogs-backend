package blogbusiness

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

type ListItemBlogStorage interface {
	ListItem(ctx context.Context, cond map[string]interface{}) ([]blogmodel.Blog, error)
}

type listItemBlogBusiness struct {
	store ListItemBlogStorage
}

func NewListItemBlogBiz(store ListItemBlogStorage) *listItemBlogBusiness {
	return &listItemBlogBusiness{
		store: store,
	}
}

func (biz *listItemBlogBusiness) ListItem(ctx context.Context, cond map[string]interface{}) ([]blogmodel.Blog, error) {
	records, err := biz.store.ListItem(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotListEntity(blogmodel.EntityName, err)
	}

	return records, nil
}
