package blogbusiness

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
)

// interface at storage of blog
type BlogStorage interface {
	GetBlog(ctx context.Context, cond map[string]interface{}) (*blogmodel.Blog, error)
}

type blogBusiness struct {
	store BlogStorage
}

func NewBlogBiz(store BlogStorage) *blogBusiness {
	return &blogBusiness{
		store: store,
	}
}

func (biz *blogBusiness) GetBlog(ctx context.Context, id int) (*blogmodel.Blog, error) {

	blog, err := biz.store.GetBlog(ctx, map[string]interface{}{"id": id, "deleted": false})

	if err != nil {
		return nil, common.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	return blog, nil
}
