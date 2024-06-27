package blogcategorybusiness

import (
	"blogs/internal/common"
	buildtree "blogs/internal/helpers/build_tree"
	blogcategorymodel "blogs/internal/model/blogcategory"
	"context"
	"strings"
)

type UpdateBlogCategoryStorage interface {
	GetBlogCategory(ctx context.Context, cond map[string]interface{}) (*blogcategorymodel.BlogCategory, error)
	ListItem(ctx context.Context, cond map[string]interface{}) ([]blogcategorymodel.BlogCategory, error)
	UpdateBlogCategory(ctx context.Context, cond map[string]interface{}, data blogcategorymodel.UpdateBlogCategory) error
}

type updateBlogCategoryBusiness struct {
	store UpdateBlogCategoryStorage
}

func NewUpdateBlogCategoryBiz(store UpdateBlogCategoryStorage) *updateBlogCategoryBusiness {
	return &updateBlogCategoryBusiness{
		store: store,
	}
}

func (biz *updateBlogCategoryBusiness) UpdateBlogCategory(ctx context.Context,
	id int, data blogcategorymodel.UpdateBlogCategory) error {

	record, err := biz.store.GetBlogCategory(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(blogcategorymodel.EntityName, err)
	}

	if err := biz.store.UpdateBlogCategory(ctx, map[string]interface{}{"id": record.Id}, data); err != nil {
		return common.ErrCannotUpdateEntity(blogcategorymodel.EntityName, err)
	}

	if strings.EqualFold(string(data.Status), string(common.StatusInactive)) {
		records, err := biz.store.ListItem(ctx, map[string]interface{}{"deleted": false})

		if err != nil {
			return common.ErrCannotListEntity(blogcategorymodel.EntityName, err)
		}

		build := buildtree.NewBuildCategoryTree(records)

		result := build.FindAllChildrenByName(record.Name)

		for _, categoryId := range result {

			var dataStatus blogcategorymodel.UpdateBlogCategory

			dataStatus.Status = common.StatusInactive

			if err := biz.store.UpdateBlogCategory(ctx, map[string]interface{}{"id": categoryId}, dataStatus); err != nil {
				return common.ErrCannotUpdateEntity(blogcategorymodel.EntityName, err)
			}
		}

	}

	return nil
}
