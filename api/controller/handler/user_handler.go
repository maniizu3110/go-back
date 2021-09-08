package handler

import (
	"net/http"
	"simplebank/api/controller/service"
	"simplebank/api/sqlc"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AssignUserHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := service.NewUserService(store)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", CreateUserHandler)
	g.GET("/", GetUserHandler)

}

func GetUserHandler(c echo.Context) error {
	service := c.Get("Service").(service.UserService)
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

func CreateUserHandler(c echo.Context) error {
	service := c.Get("Service").(service.UserService)
	params := &sqlc.CreateUserParams{}
	c.Bind(params)
	//TODO:validation
	user, err := service.CreateUser(params)
	if err != nil {
		return err
	}
	type createUserResponse struct {
		Username          string    `json:"username"`
		FullName          string    `json:"full_name"`
		Email             string    `json:"email"`
		PasswordChangedAt time.Time `json:"password_changed_at"`
		CreatedAt         time.Time `json:"created_at"`
	}
	res := createUserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusOK, res)
}
