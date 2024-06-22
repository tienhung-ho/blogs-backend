package blogcategorymodel

import "blogs/internal/common"

type CreationBlogCategory struct {
	Id             int           `json:"id" gorm:"-"`
	Name           string        `json:"name" gorm:"column:name;not null"`
	Description    string        `json:"description" gorm:"column:description;"`
	ParentCategory string        `json:"parentcategory" gorm:"column:parentcategory;"`
	Status         common.Status `json:"status" gorm:"column:status;default:'Pending'"`
	Deleted        bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (CreationBlogCategory) TableName() string {
	return BlogCategory{}.TableName()
}
