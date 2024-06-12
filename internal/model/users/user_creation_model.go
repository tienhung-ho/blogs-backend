package usersmodel

import (
	"time"
)

type UserCreation struct {
	Id        int       `json:"-" gorm:"column:id;"`
	Username  string    `json:"username" gorm:"column:username;"`
	Email     string    `json:"email" gorm:"column:email;"`
	Password  string    `json:"password" gorm:"column:password;"`
	Full_name string    `json:"full_name" gorm:"column:full_name;"`
	Birthdate time.Time `json:"birthdate" gorm:"column:birthdate;"`
	Gender    *Gender   `json:"gender" gorm:"column:gender;"`
	// Status    string    `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
}

func (UserCreation) TableName() string {
	return Users{}.TableName()
}
