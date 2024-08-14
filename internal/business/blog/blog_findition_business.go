package blogbusiness

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
	"log"
)

// BlogStorage interface at storage of blog
type BlogStorage interface {
	GetBlog(ctx context.Context, cond map[string]interface{}) (*blogmodel.Blog, error)
	CreateBlog(ctx context.Context, data *blogmodel.BlogCreation) (int, error)
}

type blogBusiness struct {
	store    BlogStorage
	rdbStore BlogStorage
}

func NewBlogBiz(store BlogStorage, rdb BlogStorage) *blogBusiness {
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
		blogCreation := blogmodel.ToBlogCreation(blog)
		_, err = biz.rdbStore.CreateBlog(ctx, blogCreation)
		if err != nil {
			return nil, err
		}
	}

	return blog, nil
}
