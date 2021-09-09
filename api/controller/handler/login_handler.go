package handler

import (
	"net/http"
	"simplebank/api/controller/services"
	"simplebank/api/sqlc"

	"github.com/labstack/echo/v4"
)

func AssignLoginHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := services.NewUserService(store)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.GET("/login", LoginUserHandler)

}

func LoginUserHandler(c echo.Context)error{
	service := c.Get("Service").(services.UserService)
	params := &services.LoginUserRequest{}
	c.Bind(params)
	res,err := service.LoginUser(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,res)

}
