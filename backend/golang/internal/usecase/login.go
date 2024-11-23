package usecase

import (
	"errors"
	"myapp/entity"
	"myapp/internal/repository"
	"myapp/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUsecase interface {
	SignIn(loginRequest model.SignInRequest) (entity.User, error)
	SignUp(signUpRequest model.SignUpRequest) error
}

type loginUsecase struct {
	userRepo repository.UserRepository
}

func NewLoginUsecase(userRepo repository.UserRepository) LoginUsecase {
	return &loginUsecase{userRepo}
}

func (u *loginUsecase) SignIn(loginRequest model.SignInRequest) (entity.User, error) {
	user, err := u.userRepo.GetUserByUserID(loginRequest.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, errors.New("uesr not found")
		} else {
			return entity.User{}, errors.New("internal server error")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return entity.User{}, errors.New("password is incorrect")
		} else {
			return entity.User{}, errors.New("internal server error")
		}
	}

	return user, nil
}

func (u *loginUsecase) SignUp(signUpRequest model.SignUpRequest) error {

	if _, err := u.userRepo.GetUserByUserID(signUpRequest.UserID); err == nil {
		return errors.New("user already exists")
	}

	if _, err := u.userRepo.GetUserByEmail(signUpRequest.Email); err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("internal server error")
	}

	user := entity.User{
		UserID:   signUpRequest.UserID,
		Email:    signUpRequest.Email,
		Password: string(hashedPassword),
		Role:     "02",
		Job:      signUpRequest.Job,
	}

	if err = u.userRepo.CreateUser(user); err != nil {
		return errors.New("internal server error")
	}

	return nil
}
