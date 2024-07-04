package blogmodel

import "blogs/internal/common"

type BlogCreation struct {
	Id          int           `json:"-" gorm:"column:id"`
	Title       string        `json:"title" gorm:"column:title;not null"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Content     string        `json:"content" gorm:"column:content;type:text"`
	AuthorID    int           `json:"authorid" gorm:"column:author_id;not null"`
	Category    string        `json:"category" gorm:"column:category;"`
	Status      common.Status `json:"status" gorm:"column:status;default:'Pending'"`
}

func (BlogCreation) TableName() string {
	return Blog{}.TableName()
}
