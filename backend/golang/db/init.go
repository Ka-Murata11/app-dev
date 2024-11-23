package db

import "myapp/entity"

func Migrate() {
	db, err := Init()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
}
