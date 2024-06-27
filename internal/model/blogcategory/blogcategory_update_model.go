package blogcategorymodel

import "blogs/internal/common"

type UpdateBlogCategory struct {
	Name           string        `json:"name" gorm:"column:name;not null"`
	Description    string        `json:"description" gorm:"column:description;"`
	ParentCategory string        `json:"parentcategory" gorm:"column:parentcategory;"`
	Status         common.Status `json:"status" gorm:"column:status;default:'Pending'"`
}

func (UpdateBlogCategory) TableName() string {
	return BlogCategory{}.TableName()
}
