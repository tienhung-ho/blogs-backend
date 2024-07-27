package permissionmodel

import "blogs/internal/common"

const (
	EntityName = "permissions"
)

type Permission struct {
	Id          int           `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string        `json:"name" gorm:"column:name;not null;unique"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Status      common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
}

func (Permission) TableName() string {
	return "permissions"
}
