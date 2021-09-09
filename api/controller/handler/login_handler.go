package handler

import (
	"fmt"
	"net/http"
	"simplebank/api/controller/services"
	"simplebank/api/sqlc"
	"simplebank/api/util"
	"simplebank/token"

	"github.com/labstack/echo/v4"
)

func AssignLoginHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			s := services.NewLoginService(store,tk,conf)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", LoginUserHandler)

}

func LoginUserHandler(c echo.Context)error{
	service := c.Get("Service").(services.LoginService)
	params := &services.LoginUserRequest{}
	c.Bind(params)
	res,err := service.LoginUser(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,res)

}
