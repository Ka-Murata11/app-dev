package usecase_test

import (
	"errors"
	"myapp/infrastructure/entity"
	"myapp/internal/model"
	"myapp/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MockPasswordUtil struct {
	mock.Mock
}

func (m *MockPasswordUtil) CreateHashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordUtil) CompareHashPassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

func TestLoginUsecase_SignIn(t *testing.T) {
	t.Run("ユーザーが存在しない", func(t *testing.T) {
		loginRequest := model.SignInRequest{
			UserID:   "test",
			Password: "password",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", loginRequest.UserID).Return(entity.User{}, gorm.ErrRecordNotFound)

		usecase := usecase.NewLoginUsecase(mockRepo, nil)

		_, err := usecase.SignIn(loginRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "user not found", err.Error())
		}
	})

	t.Run("リポジトリ内部エラー", func(t *testing.T) {
		loginRequest := model.SignInRequest{
			UserID:   "test",
			Password: "password",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", loginRequest.UserID).Return(entity.User{}, errors.New("internal server error"))

		usecase := usecase.NewLoginUsecase(mockRepo, nil)

		_, err := usecase.SignIn(loginRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "internal server error", err.Error())
		}
	})

	t.Run("パスワードが一致しない", func(t *testing.T) {
		loginRequest := model.SignInRequest{
			UserID:   "test",
			Password: "password",
		}

		acutalPassword := "notMatchedPassword"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(acutalPassword), bcrypt.DefaultCost)
		assert.NoError(t, err)
		user := entity.User{
			UserID:   "test",
			Password: string(hashedPassword),
			Email:    "test@test.com",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", loginRequest.UserID).Return(user, nil)

		mockPasswordUtil := new(MockPasswordUtil)
		mockPasswordUtil.On("CompareHashPassword", user.Password, loginRequest.Password).Return(bcrypt.ErrMismatchedHashAndPassword)

		usecase := usecase.NewLoginUsecase(mockRepo, mockPasswordUtil)

		_, err = usecase.SignIn(loginRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "password is incorrect", err.Error())
		}
	})

	t.Run("不正なハッシュ値によるエラー", func(t *testing.T) {
		loginRequest := model.SignInRequest{
			UserID:   "test",
			Password: "password",
		}

		// DBに保存されているパスワードがハッシュ化されていない
		invalidHashedPasswordPassword := "notHashPassword"
		user := entity.User{
			UserID:   "test",
			Password: string(invalidHashedPasswordPassword),
			Email:    "test@test.com",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", loginRequest.UserID).Return(user, nil)

		mockPasswordUtil := new(MockPasswordUtil)
		mockPasswordUtil.On("CompareHashPassword", user.Password, loginRequest.Password).Return(errors.New("internal server error"))

		usecase := usecase.NewLoginUsecase(mockRepo, mockPasswordUtil)

		_, err := usecase.SignIn(loginRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "internal server error", err.Error())
		}
	})

	t.Run("成功", func(t *testing.T) {
		loginRequest := model.SignInRequest{
			UserID:   "test",
			Password: "password",
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginRequest.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		expected := entity.User{
			UserID:   "test",
			Password: string(hashedPassword),
			Email:    "test@test.com",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", loginRequest.UserID).Return(expected, nil)

		mockPasswordUtil := new(MockPasswordUtil)
		mockPasswordUtil.On("CompareHashPassword", expected.Password, loginRequest.Password).Return(nil)

		usecase := usecase.NewLoginUsecase(mockRepo, mockPasswordUtil)

		actual, err := usecase.SignIn(loginRequest)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
}

func TestLogin_SignUp(t *testing.T) {
	t.Run("ユーザIDが重複している", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{
			UserID: "test",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", signUpRequest.UserID).Return(entity.User{}, nil)

		usecase := usecase.NewLoginUsecase(mockRepo, nil)

		err := usecase.SignUp(signUpRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "user already exists", err.Error())
		}
	})

	t.Run("メールアドレスが重複している", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{
			UserID: "test",
			Email:  "test@test.com",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", signUpRequest.UserID).Return(entity.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("GetUserByEmail", signUpRequest.Email).Return(entity.User{}, nil)
		usecase := usecase.NewLoginUsecase(mockRepo, nil)

		err := usecase.SignUp(signUpRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "user already exists", err.Error())
		}
	})

	t.Run("パスワード比較 内部エラー", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{
			UserID:   "test",
			Email:    "test@test.com",
			Password: "password",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", signUpRequest.UserID).Return(entity.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("GetUserByEmail", signUpRequest.Email).Return(entity.User{}, gorm.ErrRecordNotFound)

		mockUtil := new(MockPasswordUtil)
		mockUtil.On("CreateHashPassword", signUpRequest.Password).Return("", errors.New("internal server error"))

		usecase := usecase.NewLoginUsecase(mockRepo, mockUtil)

		err := usecase.SignUp(signUpRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "internal server error", err.Error())
		}
	})

	t.Run("リポジトリ内部エラー", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{
			UserID:   "test",
			Email:    "test@test.com",
			Password: "password",
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		user := entity.User{
			UserID:   signUpRequest.UserID,
			Email:    signUpRequest.Email,
			Password: string(hashedPassword),
			Role:     "02",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", signUpRequest.UserID).Return(entity.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("GetUserByEmail", signUpRequest.Email).Return(entity.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("CreateUser", user).Return(errors.New("internal server error"))

		mockUtil := new(MockPasswordUtil)
		mockUtil.On("CreateHashPassword", signUpRequest.Password).Return(string(hashedPassword), nil)

		usecase := usecase.NewLoginUsecase(mockRepo, mockUtil)

		err = usecase.SignUp(signUpRequest)
		if assert.Error(t, err) {
			assert.Equal(t, "internal server error", err.Error())
		}
	})

	t.Run("成功", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{
			UserID:   "test",
			Email:    "test@test.com",
			Password: "password",
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		user := entity.User{
			UserID:   signUpRequest.UserID,
			Email:    signUpRequest.Email,
			Password: string(hashedPassword),
			Role:     "02",
		}

		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUserByUserID", signUpRequest.UserID).Return(entity.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("GetUserByEmail", signUpRequest.Email).Return(entity.User{}, gorm.ErrRecordNotFound)
		mockRepo.On("CreateUser", user).Return(nil)

		mockUtil := new(MockPasswordUtil)
		mockUtil.On("CreateHashPassword", signUpRequest.Password).Return(string(hashedPassword), nil)

		usecase := usecase.NewLoginUsecase(mockRepo, mockUtil)

		err = usecase.SignUp(signUpRequest)
		assert.NoError(t, err)
	})
}
