package handler

import (
	"net/http"
	"simplebank/api/controller/service"
	"simplebank/api/sqlc"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AssignAccountHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := service.NewAccountService(store)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", CreateAccountHandler)
	g.GET("/:id", GetAccountHandler)
	g.GET("/", GetListAccountHandler)
	g.PUT("/",UpdateAccountHandler)
	g.DELETE("/:id",DeleteAccountHandler)

}

func GetAccountHandler(c echo.Context) error {
	service := c.Get("Service").(service.AccountService)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	account, err := service.GetAccount(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, account)
}

func CreateAccountHandler(c echo.Context) error {
	service := c.Get("Service").(service.AccountService)
	params := &sqlc.CreateAccountParams{}
	c.Bind(params)
	//TODO:validation
	account, err := service.CreateAccount(params)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, account)
}

func GetListAccountHandler(c echo.Context) error {
	service := c.Get("Service").(service.AccountService)
	params := &sqlc.ListAccountsParams{}
	c.Bind(params)
	accounts,err := service.GetListAccount(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,accounts)
}

func UpdateAccountHandler(c echo.Context) error {
	service := c.Get("Service").(service.AccountService)
	params := &sqlc.UpdateAccountParams{}
	c.Bind(params)
	account,err := service.UpdateAccount(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,account)
}
func DeleteAccountHandler(c echo.Context) error {
	service := c.Get("Service").(service.AccountService)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = service.DeleteAccount(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK,"アカウントの削除が完了しました")
}


