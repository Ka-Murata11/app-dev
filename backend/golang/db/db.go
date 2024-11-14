package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	return nil, err
	// }

	// dsn := os.Getenv("DEV_DB_DSN")
	// fmt.Println(dsn + "aa")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("user:password@tcp(db:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
