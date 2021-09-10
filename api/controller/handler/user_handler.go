package handler

import (
	"net/http"
	"simplebank/api/controller/services"
	"simplebank/api/sqlc"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AssignUserHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := services.NewUserService(store)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.GET("/", GetUserHandler)
}

func GetUserHandler(c echo.Context) error {
	service := c.Get("Service").(services.UserService)
	type getUserParams struct {
		Username string
	}
	params := &getUserParams{}
	//TODO:エラーハンドリング
	c.Bind(params)
	User, err := service.GetUser(params.Username)
	if err != nil {
		logrus.Info(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, User)
}

