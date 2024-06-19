package usersmodel

import (
	"blogs/internal/common"
	"time"
)

type UserCreation struct {
	Id        int           `json:"-" gorm:"column:id;"`
	Username  string        `json:"username" gorm:"column:username;"`
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Full_name string        `json:"full_name" gorm:"column:full_name;"`
	Birthdate time.Time     `json:"birthdate" gorm:"column:birthdate;"`
	Gender    common.Gender `json:"gender" gorm:"column:gender;"`
	Status    common.Status `json:"status" gorm:"column:status;default:'Pending'"`
	Deleted   bool          `json:"deleted" gorm:"column:deleted;default:false"`
}

func (UserCreation) TableName() string {
	return Users{}.TableName()
}
