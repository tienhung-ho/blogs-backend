package buildtree

import blogcategorymodel "blogs/internal/model/blogcategory"

// FindAllChildrenByID finds and returns the names of all children and their descendants of a category with a given ID.
func FindAllChildrenByID(categories []blogcategorymodel.ListBlogCategory, id int) []int {
	var result []int

	for _, category := range categories {
		if category.Id == id {
			collectChildrenNames(&result, category.Child)
			return result
		}

		childResult := FindAllChildrenByID(category.Child, id)

		if len(childResult) > 0 {
			return childResult
		}
	}
	return result
}

// collectChildrenNames collects the names of all children and their descendants recursively.
func collectChildrenNames(result *[]int, children []blogcategorymodel.ListBlogCategory) {
	for _, child := range children {
		*result = append(*result, child.Id)
		collectChildrenNames(result, child.Child)
	}
}
