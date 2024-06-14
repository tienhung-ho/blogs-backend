package authmodel

import (
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	"time"
)

type UserRegister struct {
	Id        int           `json:"-" gorm:"column:id;"`
	Username  string        `json:"username" gorm:"column:username;"`
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Full_name string        `json:"full_name" gorm:"column:full_name;"`
	Birthdate time.Time     `json:"birthdate" gorm:"column:birthdate;"`
	Gender    common.Gender `json:"gender" gorm:"column:gender;"`
	Status    common.Status `json:"status" gorm:"column:status;default:'Pending'"`
}

func (UserRegister) TableName() string {
	return usersmodel.Users{}.TableName()
}

func (user UserRegister) DoRegister(hashPass string) *usersmodel.UserCreation {
	return &usersmodel.UserCreation{
		Username:  user.Username,
		Email:     user.Email,
		Password:  hashPass,
		Full_name: user.Full_name,
		Birthdate: user.Birthdate,
		Gender:    user.Gender,
		Status:    user.Status,
	}
}
