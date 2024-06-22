package blogcategorybusiness

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

type CreationBlogCategoryStorage interface {
	GetBlogCategory(ctx context.Context, cond map[string]interface{}) (*blogcategorymodel.BlogCategory, error)
	CreateBlogCategory(ctx context.Context, data *blogcategorymodel.CreationBlogCategory) (int, error)
}

type creationBlogCategoryBusiness struct {
	store CreationBlogCategoryStorage
}

func NewCreationBlogCategoryBiz(store CreationBlogCategoryStorage) *creationBlogCategoryBusiness {
	return &creationBlogCategoryBusiness{
		store: store,
	}
}

func (biz *creationBlogCategoryBusiness) CreateBlogCategory(ctx context.Context, createData *blogcategorymodel.CreationBlogCategory) (int, error) {
	record, err := biz.store.GetBlogCategory(ctx, map[string]interface{}{"name": createData.Name})

	if record != nil {
		return 0, common.ErrRecordExist(blogcategorymodel.EntityName, err)
	}

	dataId, err := biz.store.CreateBlogCategory(ctx, createData)

	if err != nil {
		return 0, common.ErrCannotCreateEntity(blogcategorymodel.EntityName, err)
	}

	return dataId, nil
}
