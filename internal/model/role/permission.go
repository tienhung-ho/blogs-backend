package rolemodel

import (
	"blogs/internal/common"
)

const (
	PermissionEntityName = "Permission"
)

type Permission struct {
	Id          int           `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string        `json:"name" gorm:"column:name;not null;unique"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Roles       []Role        `gorm:"many2many:role_permissions;"`
	Status      common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
}

func (Permission) TableName() string {
	return "permissions"
}
