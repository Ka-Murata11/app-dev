package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("user:password@tcp(db:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
