package middleware

import (
	db "simplebank/db/sqlc"

	"github.com/labstack/echo/v4"
)


func SetStore(store *db.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO:storeを変数でもつ
			c.Set("store", store)
			return nil
		}
	}
}