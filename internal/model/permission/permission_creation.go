package permissionmodel

import "blogs/internal/common"

type PermissionCreation struct {
	Id          int           `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string        `json:"name" gorm:"column:name;not null;unique"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Status      common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
	Deleted     bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (PermissionCreation) TableName() string {
	return Permission{}.TableName()
}
