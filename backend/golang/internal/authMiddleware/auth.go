package authMiddleware

import (
	"myapp/auth"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, err := auth.ParseToken(cookie.Value)
		if err != nil {
			return echo.ErrUnauthorized
		}

		// 新しいトークンを生成
		newToken, err := auth.CreateToken(claims.UserID, claims.Role)
		if err != nil {
			return echo.ErrInternalServerError
		}

		// 新しいトークンをCookieに設定
		auth.SetCookie(c, newToken)

		return next(c)
	}
}
