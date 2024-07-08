package blogbusiness

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
	"log"
)

type BlogDeletionStorage interface {
	GetBlog(ctx context.Context, cond map[string]interface{}) (*blogmodel.Blog, error)
	DeleteBlog(ctx context.Context, cond map[string]interface{}) error
}

type blogDeletionBusiness struct {
	store BlogDeletionStorage
}

func NewBlogDeletionBiz(store BlogDeletionStorage) *blogDeletionBusiness {
	return &blogDeletionBusiness{
		store: store,
	}
}

func (biz *blogDeletionBusiness) DeleteBlog(ctx context.Context, id int) error {

	record, err := biz.store.GetBlog(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	if err := biz.store.DeleteBlog(ctx, map[string]interface{}{"id": record.Id}); err != nil {
		log.Printf("%v", err)
		return common.ErrCannotDeleteEntity(blogmodel.EntityName, err)
	}

	return nil
}
