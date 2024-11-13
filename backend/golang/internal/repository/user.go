package repository

import (
	"myapp/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
