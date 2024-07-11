package rolemodel

import "time"

const (
	RoleEntityName = "role"
)

type Role struct {
	Id          int          `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string       `json:"name" gorm:"column:name;not null;unique"`
	Description string       `json:"description" gorm:"column:description;type:text"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time    `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Role) TableName() string {
	return "roles"
}
