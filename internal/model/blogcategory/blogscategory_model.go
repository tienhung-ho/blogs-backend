package blogcategorymodel

import "blogs/internal/common"

const (
	EntityName = "BlogCategory"
)

type BlogCategory struct {
	Id             int           `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string        `json:"name" gorm:"column:name;not null;index:idx_blog_categories_name"`
	Description    string        `json:"description" gorm:"column:description"`
	ParentCategory string        `json:"parentcategory" gorm:"column:parentcategory;index:idx_blog_categories_parentcategory"`
	Status         common.Status `json:"status" gorm:"column:status;type:enum('Pending','Active','Inactive');default:'Pending';index:idx_blog_categories_status"`
	Deleted        bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (BlogCategory) TableName() string {
	return "blog_categories"
}

func (b *BlogCategory) ToBlogCategoryList() *ListBlogCategory {
	return &ListBlogCategory{
		Id:             b.Id,
		Name:           b.Name,
		Description:    b.Description,
		ParentCategory: b.ParentCategory,
	}
}
