package buildtree

import (
	blogcategorymodel "blogs/internal/model/blogcategory"
)

type buildCategoryTree struct {
	data []blogcategorymodel.BlogCategory
}

func NewBuildCategoryTree(data []blogcategorymodel.BlogCategory) *buildCategoryTree {
	return &buildCategoryTree{
		data: data,
	}
}

func (b *buildCategoryTree) BuildCategoryTree(parent string) []blogcategorymodel.ListBlogCategory {
	var result []blogcategorymodel.ListBlogCategory

	for _, category := range b.data {
		if category.ParentCategory == parent {
			// Chuyển đổi BlogCategory sang ListBlogCategory
			newCategory := category.ToBlogCategoryList()
			newCategory.Child = b.BuildCategoryTree(category.Name) // Đệ quy để tìm các con
			result = append(result, *newCategory)
		}
	}

	return result
}
