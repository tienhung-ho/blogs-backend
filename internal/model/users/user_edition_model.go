package usersmodel

import (
	"blogs/internal/common"
	"time"
)

type UserEdition struct {
	Username  string        `json:"username" form:"username" gorm:"column:username;"`
	Email     string        `json:"email" form:"email" gorm:"column:email;"`
	Password  string        `json:"password" form:"password" gorm:"column:password;"`
	Full_name string        `json:"full_name" form:"full_name" gorm:"column:full_name;"`
	Birthdate time.Time     `json:"birthdate" form:"birthdate" gorm:"column:birthdate;"`
	Gender    common.Gender `json:"gender" form:"gener" gorm:"column:gender;"`
	Status    common.Status `json:"status" form:"status" gorm:"column:status;default:'Pending'"`
	Deleted   bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (UserEdition) TableName() string {
	return Users{}.TableName()
}
