package rolemodel

import (
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
)

const (
	RoleEntityName = "role"
)

type Permission struct {
	Id          int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"column:name;not null;unique;index:idx_permissions_name"`
	Description string `json:"description" gorm:"column:description;type:text"`
	//Roles       []Role        `gorm:"many2many:role_permissions;"`
	Status common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending');index:idx_permissions_status"`
}

func (Permission) TableName() string {
	return "permissions"
}

type Role struct {
	Id          int                          `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string                       `json:"name" gorm:"column:name;not null;unique;index:idx_roles_name"`
	Description string                       `json:"description" gorm:"column:description;type:text"`
	Permissions []permissionmodel.Permission `gorm:"many2many:role_permissions;"`
	Status      common.Status                `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending');index:idx_roles_status"`
}

func (Role) TableName() string {
	return "roles"
}
