package handler

import (
	"fmt"
	"net/http"
	"simplebank/api/sqlc"

	// db "simplebank/db/sqlc"

	"github.com/labstack/echo/v4"
)

func AssignAccountHandler(g *echo.Group) {
	fmt.Println("assign")
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// store := c.Get("store").(*db.Store)
			return handler(c)
		}
	})
	g.GET("/", GetAccountHandler)
}

func GetAccountHandler(c echo.Context) error {
	store := c.Get("store").(*sqlc.Store)
	account,err := store.GetAccount(c.Request().Context(),1)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, account)

}
