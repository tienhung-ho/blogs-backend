package blogcategorybusiness

import (
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

type FinditionBlogCategoryStorage interface {
	GetBlogCategory(ctx context.Context, cond map[string]interface{}) (*blogcategorymodel.BlogCategory, error)
}

type finditionBlogCategoryBiz struct {
	store FinditionBlogCategoryStorage
}

func NewFinditionBlogCategory(store FinditionBlogCategoryStorage) *finditionBlogCategoryBiz {
	return &finditionBlogCategoryBiz{
		store: store,
	}
}

func (biz *finditionBlogCategoryBiz) GetBlogCategoryByCondition(ctx context.Context, cond map[string]interface{}) (*blogcategorymodel.BlogCategory, error) {
	record, err := biz.store.GetBlogCategory(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(blogcategorymodel.EntityName, err)
	}

	return record, nil
}
