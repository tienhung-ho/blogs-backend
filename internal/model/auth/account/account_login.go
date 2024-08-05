package accountautmodel

import accountmodel "blogs/internal/model/accounts"

type AccountLogin struct {
	Username string `json:"username" gorm:"column:username;not null;unique" form:"username"`
	Password string `json:"password" gorm:"column:password;not null" form:"password"`
}

func (AccountLogin) TableName() string {
	return accountmodel.Account{}.TableName()
}
