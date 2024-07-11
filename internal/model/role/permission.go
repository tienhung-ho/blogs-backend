package rolemodel

import (
	"time"
)

const (
	PermissionEntityName = "Permission"
)

type Permission struct {
	Id          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"column:name;not null;unique"`
	Description string    `json:"description" gorm:"column:description;type:text"`
	Roles       []Role    `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Permission) TableName() string {
	return "permissions"
}
