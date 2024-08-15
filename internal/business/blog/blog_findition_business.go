package blogbusiness

import (
	"blogs/internal/common"
	cachehelper "blogs/internal/helpers/cache"
	blogmodel "blogs/internal/model/blog"
	filtermodel "blogs/internal/model/filter"
	"context"
	"log"
)

// BlogStorage interface at storage of blog
type BlogStorage interface {
	GetBlog(ctx context.Context, cond map[string]interface{}, morekeys ...string) (*blogmodel.Blog, error)
	CreateBlog(ctx context.Context, data *blogmodel.BlogCreation, morekeys ...string) (int, error)
}

type BlogCacheStorage interface {
	GetBlog(ctx context.Context, cond map[string]interface{}, morekeys ...string) (*blogmodel.Blog, error)
	CreateBlog(ctx context.Context, data interface{}, morekeys ...string) (int, error)
}

type blogBusiness struct {
	store    BlogStorage
	rdbStore BlogCacheStorage
}

func NewBlogBiz(store BlogStorage, rdb BlogCacheStorage) *blogBusiness {
	return &blogBusiness{
		store:    store,
		rdbStore: rdb,
	}
}

func (biz *blogBusiness) GetBlog(ctx context.Context, id int) (*blogmodel.Blog, error) {

	blog, err := biz.rdbStore.GetBlog(ctx, map[string]interface{}{"id": id})

	if err != nil {
		log.Printf("%v", err)
		return nil, common.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	//If blog is found in cache, return it
	if blog != nil {
		return blog, nil
	}
	blog, err = biz.store.GetBlog(ctx, map[string]interface{}{"id": id, "deleted": false})

	if err != nil {
		return nil, common.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	//Save the blog to cache for future requests
	if blog != nil {
		var paging common.Paging

		paging.Process()
		key := cachehelper.GenerateCacheKey(blogmodel.EntityName, map[string]interface{}{"id": id}, paging, filtermodel.Filter{})
		blogCreation := blogmodel.ToBlogCreation(blog)
		_, err = biz.rdbStore.CreateBlog(ctx, blogCreation, key)
		if err != nil {
			return nil, err
		}
	}

	return blog, nil
}
