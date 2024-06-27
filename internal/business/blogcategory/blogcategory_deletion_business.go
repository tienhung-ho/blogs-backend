package blogcategorybusiness

import (
	"blogs/internal/common"
	buildtree "blogs/internal/helpers/build_tree"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
)

type DeletionBlogCategoryStorage interface {
	GetBlogCategory(ctx context.Context, cond map[string]interface{}) (*blogcategorymodel.BlogCategory, error)
	ListItem(ctx context.Context, cond map[string]interface{}) ([]blogcategorymodel.BlogCategory, error)
	DeleteBlogCategory(ctx context.Context, cond map[string]interface{}) error
}

type deletionBlogCategoryBusiness struct {
	store DeletionBlogCategoryStorage
}

func NewDeletionBlogCategoryBusiness(store DeletionBlogCategoryStorage) *deletionBlogCategoryBusiness {
	return &deletionBlogCategoryBusiness{
		store: store,
	}
}

func (biz *deletionBlogCategoryBusiness) DeleteBlogCategory(ctx context.Context, id int) error {

	record, err := biz.store.GetBlogCategory(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(blogcategorymodel.EntityName, err)
	}

	records, err := biz.store.ListItem(ctx, map[string]interface{}{"deleted": false})

	if err != nil {
		return common.ErrCannotListEntity(blogcategorymodel.EntityName, err)
	}

	if err := biz.store.DeleteBlogCategory(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(blogcategorymodel.EntityName, err)
	}

	build := buildtree.NewBuildCategoryTree(records)

	result := build.FindAllChildrenByName(record.Name)

	for _, categoryId := range result {
		if err := biz.store.DeleteBlogCategory(ctx, map[string]interface{}{"id": categoryId}); err != nil {
			return common.ErrCannotDeleteEntity(blogcategorymodel.EntityName, err)
		}
	}

	return nil
}
