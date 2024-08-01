package accountsstogare

import "gorm.io/gorm"

type mysqlStorage struct {
	db *gorm.DB
}

func NewMysqlStorage(db *gorm.DB) *mysqlStorage {
	return &mysqlStorage{
		db: db,
	}
}
