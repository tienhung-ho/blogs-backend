package authmodel

import usersmodel "blogs/internal/model/users"

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return usersmodel.Users{}.TableName()
}
