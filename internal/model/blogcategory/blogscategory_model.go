package blogcategorymodel

import "blogs/internal/common"

const (
	EntityName = "BlogCategory"
)

type BlogCategory struct {
	Id             int           `json:"-" gorm:"column:id;"`
	Name           string        `json:"name" gorm:"column:name;not null"`
	Description    string        `json:"description" gorm:"column:description;"`
	ParentCategory string        `json:"parentcategory" gorm:"column:parentcategory;"`
	Status         common.Status `json:"status" gorm:"column:status;default:'Pending'"`
	Deleted        bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (BlogCategory) TableName() string {
	return "blog_categories"
}
