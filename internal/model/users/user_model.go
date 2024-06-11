package usersmodel

import (
	"database/sql"
	"errors"
)

const (
	EntityName = "User"
)

var (
	ErrUserNameBlank = errors.New("cannot leave username blank")
)

type Users struct {
	Username  string       `json:"username" gorm:"column:username;"`
	Email     string       `json:"email" gorm:"column:email;"`
	Password  string       `json:"password" gorm:"column:password;"`
	Full_name string       `json:"full_name" gorm:"column:full_name;"`
	Birthdate sql.NullTime `json:"birthdate" gorm:"column:birthdate;type:date"`
	Gender    string       `json:"gender" gorm:"column:gender;"`
}

func (Users) TableName() string {
	return "users"
}
