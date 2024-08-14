package blogbusiness

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

type UpdateBlogStorage interface {
	GetBlog(ctx context.Context, cond map[string]interface{}) (*blogmodel.Blog, error)
	UpdateBlog(ctx context.Context, cond map[string]interface{}, data *blogmodel.BlogUpdate) error
}

type UpdateBlogCacheStorage interface {
	DeleteBlog(ctx context.Context, cond map[string]interface{}) error
}

type updateBlogBusiness struct {
	store    UpdateBlogStorage
	rdbStore UpdateBlogCacheStorage
}

func NewUpdateBlogBiz(store UpdateBlogStorage, rdbStore UpdateBlogCacheStorage) *updateBlogBusiness {
	return &updateBlogBusiness{
		store:    store,
		rdbStore: rdbStore,
	}
}

func (biz *updateBlogBusiness) UpdateBlog(ctx context.Context, cond map[string]interface{}, data *blogmodel.BlogUpdate) error {
	record, err := biz.store.GetBlog(ctx, cond)

	if err != nil {
		return common.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	if record.Deleted {
		return common.ErrEntityDeleted(blogmodel.EntityName, err)
	}

	if err := biz.store.UpdateBlog(ctx, cond, data); err != nil {
		return common.ErrCannotUpdateEntity(blogmodel.EntityName, err)
	}

	if err := biz.rdbStore.DeleteBlog(ctx, cond); err != nil {

		return common.ErrCannotDeleteEntity(blogmodel.EntityName, err)
	}

	return nil
}
