package accountmodel

import "blogs/internal/common"

const (
	EntityName = "accounts"
)

type Account struct {
	ID       int            `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Username string         `json:"username" gorm:"column:username;not null;unique;uniqueIndex:idx_accounts_username" form:"username"`
	Email    string         `json:"email" gorm:"column:email;not null;unique;uniqueIndex:idx_accounts_email" form:"email"`
	Password string         `json:"password" gorm:"column:password;not null" form:"password"`
	FullName string         `json:"full_name" gorm:"column:full_name" form:"full_name"`
	Image    string         `json:"-" gorm:"column:image"`
	RoleID   int            `json:"role_id" gorm:"column:role_id;index:idx_accounts_role_id" form:"role_id"`
	Gender   *common.Gender `json:"gender" gorm:"column:gender;type:enum('Male','Female','Other')" form:"gender"`
	Status   *common.Status `json:"status" gorm:"column:status;type:enum('Pending','Active','Inactive');default:'Pending';index:idx_accounts_status" form:"status"`
	Deleted  bool           `json:"deleted" gorm:"column:deleted;default:false" form:"deleted"`
}

// TableName sets the insert table name for this struct type
func (Account) TableName() string {
	return "accounts"
}

type SimpleAccount struct {
	ID       int            `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Username string         `json:"username" gorm:"column:username;not null;unique;uniqueIndex:idx_accounts_username" form:"username"`
	Email    string         `json:"email" gorm:"column:email;not null;unique;uniqueIndex:idx_accounts_email" form:"email"`
	FullName string         `json:"full_name" gorm:"column:full_name" form:"full_name"`
	Image    string         `json:"-" gorm:"column:image"`
	RoleID   int            `json:"role_id" gorm:"column:role_id" form:"role_id"`
	Gender   *common.Gender `json:"gender" gorm:"column:gender;type:enum('Male','Female','Other')" form:"gender"`
}

func (acc *Account) ToSimpleAccount() *SimpleAccount {
	return &SimpleAccount{
		ID:       acc.ID,
		Username: acc.Username,
		Email:    acc.Email,
		FullName: acc.FullName,
		Image:    acc.Image,
		RoleID:   acc.RoleID,
		Gender:   acc.Gender,
	}
}
