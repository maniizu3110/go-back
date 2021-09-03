package handler

import (
	"fmt"
	"net/http"
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
	// store := c.Get("store").(db.Store)
	// fmt.Printf("store:%v",store)

	// fmt.Println("handler")
	return c.JSON(http.StatusOK, "ここまできてるよ")

}
