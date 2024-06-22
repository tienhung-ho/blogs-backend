package blogcategorymodel

type ListBlogCategory struct {
	Name           string             `json:"name" gorm:"column:name;not null"`
	Description    string             `json:"description" gorm:"column:description;"`
	ParentCategory string             `json:"parentcategory" gorm:"column:parentcategory;"`
	Child          []ListBlogCategory `json:"child,omitempty"`
}

func (ListBlogCategory) TableName() string {
	return BlogCategory{}.TableName()
}
