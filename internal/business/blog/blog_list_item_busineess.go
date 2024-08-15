package blogbusiness

import (
	"blogs/internal/common"
	cachehelper "blogs/internal/helpers/cache"
	blogmodel "blogs/internal/model/blog"
	filtermodel "blogs/internal/model/filter"
	"context"
)

type ListItemBlogStorage interface {
	ListItem(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekeys ...string) ([]blogmodel.Blog, error)
	CreateBlog(ctx context.Context, blog *blogmodel.BlogCreation, morekeys ...string) (int, error)
}

type CacheStorage interface {
	ListItem(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekeys ...string) ([]blogmodel.Blog, error)
	CreateBlog(ctx context.Context, blog interface{}, morekeys ...string) (int, error)
}

type listItemBlogBusiness struct {
	store    ListItemBlogStorage
	rdbStore CacheStorage
}

func NewListItemBlogBiz(store ListItemBlogStorage, rdbStore CacheStorage) *listItemBlogBusiness {
	return &listItemBlogBusiness{
		store:    store,
		rdbStore: rdbStore,
	}
}

func (biz *listItemBlogBusiness) ListItem(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter) ([]blogmodel.Blog, error) {

	// Kiểm tra cache trước
	cachedData, err := biz.rdbStore.ListItem(ctx, cond, paging, filter)
	if err != nil {
		return nil, err
	}
	if cachedData != nil {
		return cachedData, nil
	}

	records, err := biz.store.ListItem(ctx, cond, paging, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(blogmodel.EntityName, err)
	}

	if len(records) != 0 {
		key := cachehelper.GenerateCacheKey(blogmodel.EntityName, cond, *paging, *filter)
		_, err := biz.rdbStore.CreateBlog(ctx, records, key)

		if err != nil {
			return nil, common.ErrCannotListEntity(blogmodel.EntityName, err)
		}
	}

	return records, nil
}
