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

func (b *buildCategoryTree) FindAllChildrenByName(parent string) []int {
	var result []int
	b.findChildren(parent, &result)
	return result
}

func (b *buildCategoryTree) findChildren(parent string, result *[]int) {
	for _, category := range b.data {
		if category.ParentCategory == parent {
			*result = append(*result, category.Id)
			b.findChildren(category.Name, result) // Đệ quy tìm các danh mục con của danh mục hiện tại
		}
	}
}
