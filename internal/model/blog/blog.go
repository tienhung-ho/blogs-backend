package blogmodel

import "blogs/internal/common"

const (
	EntityName = "blog"
)

type Blog struct {
	Id          int           `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Title       string        `json:"title" gorm:"column:title;not null;index:idx_blog_title"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Content     string        `json:"content" gorm:"column:content;type:text"`
	AuthorID    int           `json:"authorid" gorm:"column:author_id;not null;index:idx_blog_author_id"`
	Category    string        `json:"category" gorm:"column:category;index:idx_blog_category"`
	Status      common.Status `json:"status" gorm:"column:status;type:enum('Pending','Active','Inactive');default:'Pending';index:idx_blog_status"`
	Deleted     bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (Blog) TableName() string {
	return "blogs"
}

func ToBlogCreation(b *Blog) *BlogCreation {
	return &BlogCreation{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		Content:     b.Content,
		AuthorID:    b.AuthorID,
		Category:    b.Category,
		Status:      b.Status,
	}
}
