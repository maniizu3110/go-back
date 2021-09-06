package middleware

import (
	"simplebank/api/sqlc"

	"github.com/labstack/echo/v4"
)

func SetStore(store sqlc.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO:storeを変数で管理する
			c.Set("store", store)
			return next(c)
		}
	}
}
