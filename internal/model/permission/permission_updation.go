package permissionmodel

import "blogs/internal/common"

type PermissionUpdation struct {
	Id          int           `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string        `json:"name" gorm:"column:name;not null;unique"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Status      common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
}

func (PermissionUpdation) TableName() string {
	return Permission{}.TableName()
}
