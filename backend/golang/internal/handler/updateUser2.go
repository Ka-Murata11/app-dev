package handler

import (
	"myapp/db"
	"myapp/infrastructure/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateUserInf struct {
	UserID string `query:"user_id" validate:"required"`
	Email  string `query:"email" validate:"email"`
	Role   string `query:"role"`
	Job    string `query:"job"`
}

func UpdateUser(c echo.Context) error {
	var updateUserInf UpdateUserInf
	if err := c.Bind(&updateUserInf); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&updateUserInf); err != nil {
		return echo.ErrBadRequest
	}

	db, err := db.Init()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var user entity.User
	if err := db.Where("user_id = ?", updateUserInf.UserID).First(&user).Error; err != nil {
		return echo.ErrBadRequest
	}

	user.Email = updateUserInf.Email
	user.Job = updateUserInf.Job
	user.Role = updateUserInf.Role

	// if err := db.Model(&user).Where("user_id = ?", updateUserInf.UserID).Updates(map[string]interface{}{"email": updateUserInf.Email,
	// 	"role": updateUserInf.Role, "job": updateUserInf.Job}).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return echo.ErrBadRequest
	// 	} else {
	// 		return echo.ErrInternalServerError
	// 	}

	// }

	if err := db.Save(&user).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, struct{}{})
}
