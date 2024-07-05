package blogmodel

import "blogs/internal/common"

type BlogUpdate struct {
	Id          int           `json:"-" gorm:"column:id"`
	Title       string        `json:"title" gorm:"column:title;not null"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Content     string        `json:"content" gorm:"column:content;type:text"`
	Category    string        `json:"category" gorm:"column:category;"`
	Status      common.Status `json:"status" gorm:"column:status;default:'Pending'"`
}

func (BlogUpdate) TableName() string {
	return Blog{}.TableName()
}
