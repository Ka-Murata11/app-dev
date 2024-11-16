package repository_test

import (
	"myapp/entity"
	"myapp/internal/repository"
	"myapp/tests/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_GetUsers(t *testing.T) {
	db, err := testutil.SetupTestDB()
	assert.NoError(t, err)

	userRepo := repository.NewUserRepository(db)

	t.Run("0件取得", func(t *testing.T) {
		users, err := userRepo.GetUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 0)
	})

	t.Run("1件取得", func(t *testing.T) {
		expected := entity.User{UserID: "user1", Email: "test1@test.com", Password: "password"}
		err := db.Create(&expected).Error
		assert.NoError(t, err)

		users, err := userRepo.GetUsers()
		assert.NoError(t, err)
		if assert.Len(t, users, 1) {
			assert.Equal(t, expected, users[0])
		}

		err = db.Where("user_id = ?", expected.UserID).Delete(&entity.User{}).Error
		assert.NoError(t, err)
	})

	t.Run("2件取得", func(t *testing.T) {
		expected := []entity.User{
			{UserID: "user1", Email: "test1@test.com", Password: "password"},
			{UserID: "user2", Email: "test2@test.com", Password: "password"},
		}
		err := db.Create(&expected).Error
		assert.NoError(t, err)

		users, err := userRepo.GetUsers()
		assert.NoError(t, err)
		if assert.Len(t, users, 2) {
			assert.Equal(t, expected, users)
		}
	})
}

func TestUserRepository_GetUserByUserID(t *testing.T) {
	db, err := testutil.SetupTestDB()
	assert.NoError(t, err)

	userRepo := repository.NewUserRepository(db)

	t.Run("存在しないユーザー", func(t *testing.T) {
		user, err := userRepo.GetUserByUserID("user1")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
		assert.Equal(t, entity.User{}, user)
	})

	t.Run("存在するユーザー", func(t *testing.T) {
		expected := []entity.User{
			{UserID: "user1", Email: "test1@test.com", Password: "password"},
			{UserID: "user2", Email: "test2@test.com", Password: "password"},
		}
		err := db.Create(&expected).Error
		assert.NoError(t, err)

		user, err := userRepo.GetUserByUserID("user1")
		assert.NoError(t, err)
		assert.Equal(t, expected[0], user)
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	db, err := testutil.SetupTestDB()
	assert.NoError(t, err)

	userRepo := repository.NewUserRepository(db)

	t.Run("存在しないユーザー", func(t *testing.T) {
		user, err := userRepo.GetUserByEmail("test1@test.com")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
		assert.Equal(t, entity.User{}, user)
	})

	t.Run("存在するユーザー", func(t *testing.T) {
		expected := []entity.User{
			{UserID: "user1", Email: "test1@test.com", Password: "password"},
			{UserID: "user2", Email: "test2@test.com", Password: "password"},
		}
		err := db.Create(&expected).Error
		assert.NoError(t, err)

		user, err := userRepo.GetUserByEmail("test1@test.com")
		assert.NoError(t, err)
		assert.Equal(t, expected[0], user)
	})
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, err := testutil.SetupTestDB()
	assert.NoError(t, err)

	userRepo := repository.NewUserRepository(db)

	addUsers := []entity.User{
		{UserID: "user1", Email: "test1@test.com", Password: "password"},
		{UserID: "user2", Email: "test2@test.com", Password: "password"},
	}
	err = db.Create(&addUsers).Error
	assert.NoError(t, err)

	t.Run("成功", func(t *testing.T) {
		addUser := entity.User{UserID: "user3", Email: "test3@test.com", Password: "password"}
		err := userRepo.CreateUser(addUser)
		assert.NoError(t, err)

		expected := []entity.User{
			{UserID: "user1", Email: "test1@test.com", Password: "password"},
			{UserID: "user2", Email: "test2@test.com", Password: "password"},
			{UserID: "user3", Email: "test3@test.com", Password: "password"},
		}

		var actual []entity.User
		err = db.Find(&actual).Error
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
