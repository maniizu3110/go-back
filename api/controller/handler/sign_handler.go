package handler

import (
	"net/http"
	"simplebank/api/controller/services"
	"simplebank/api/sqlc"

	"github.com/labstack/echo/v4"
)

func AssignSignHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := services.NewSignService(store)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", CreateUserHandler)
}

func CreateUserHandler(c echo.Context) error {
	service := c.Get("Service").(services.SignService)
	params := &sqlc.CreateUserParams{}
	c.Bind(params)
	//TODO:validation
	user, err := service.CreateUser(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}