package usersmodel

import (
	"blogs/internal/common"
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
	ID        int           `json:"id" gorm:"column:id;"`
	Username  string        `json:"username" gorm:"column:username;"`
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Full_name string        `json:"full_name" gorm:"column:full_name;"`
	Birthdate sql.NullTime  `json:"birthdate" gorm:"column:birthdate;type:date"`
	Gender    common.Gender `json:"gender" gorm:"column:gender;type:enum('Male','Female','Other')"`
	Status    common.Status `json:"status" gorm:"column:status;type:enum('Active','Inactive','Pending')"`
}

func (Users) TableName() string {
	return "users"
}
