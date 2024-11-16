package repository

import (
	"myapp/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUserByUserID(userID string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) error
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
		return []entity.User{}, err
	}

	return users, nil
}

func (r *userRepository) GetUserByUserID(userID string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user entity.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			return err
		}
	}

	return nil
}
