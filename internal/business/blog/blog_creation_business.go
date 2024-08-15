package blogbusiness

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
	"errors"
)

type BlogCreationStorage interface {
	CreateBlog(ctx context.Context, data *blogmodel.BlogCreation, morekeys ...string) (int, error)
}

type blogCreationBusiness struct {
	store BlogCreationStorage
}

func NewBlogCreationBusiness(store BlogCreationStorage) *blogCreationBusiness {
	return &blogCreationBusiness{
		store: store,
	}
}

func (biz *blogCreationBusiness) CreateBlog(ctx context.Context, data *blogmodel.BlogCreation) (int, error) {

	if len(data.Content) < 1 || len(data.Title) < 1 {

		return 0, common.ErrInternal(errors.New("empty title or content"))
	}

	dataId, err := biz.store.CreateBlog(ctx, data)

	if err != nil {
		return 0, common.ErrCannotCreateEntity(blogmodel.EntityName, err)
	}

	return dataId, nil
}
