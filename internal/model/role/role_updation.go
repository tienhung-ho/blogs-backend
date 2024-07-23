package rolemodel

import "blogs/internal/common"

type RoleUpdation struct {
	Id          int           `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string        `json:"name" gorm:"column:name;not null;unique"`
	Description string        `json:"description" gorm:"column:description;type:text"`
	Status      common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
	Permissions []Permission  `gorm:"many2many:role_permissions;foreignKey:Id;joinForeignKey:RoleId;References:Id;joinReferences:PermissionId"`
}

func (RoleUpdation) TableName() string {
	return Role{}.TableName()
}
