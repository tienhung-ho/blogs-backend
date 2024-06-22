package blogcategorybusiness

import (
	"blogs/internal/common"
	buildtree "blogs/internal/helpers/build_tree"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

// get item with interface
type ListItemBlogCategoryStorage interface {
	ListItem(ctx context.Context, cond map[string]interface{}) ([]blogcategorymodel.BlogCategory, error)
}

type listItemBlogCategoryBusiness struct {
	store ListItemBlogCategoryStorage
}

func NewListItemBlogCategoryStorage(store ListItemBlogCategoryStorage) *listItemBlogCategoryBusiness {
	return &listItemBlogCategoryBusiness{
		store: store,
	}
}

func (biz *listItemBlogCategoryBusiness) ListItem(ctx context.Context, cond map[string]interface{}) ([]blogcategorymodel.ListBlogCategory, error) {

	records, err := biz.store.ListItem(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotListEntity(blogcategorymodel.EntityName, err)
	}

	build := buildtree.NewBuildCategoryTree(records)

	categoryTree := build.BuildCategoryTree("")

	return categoryTree, nil
}
