package blogcategorystorage

import "gorm.io/gorm"

type mysqlStorage struct {
	db *gorm.DB
}

func NewSqlStorage(db *gorm.DB) *mysqlStorage {
	return &mysqlStorage{db: db}
}
