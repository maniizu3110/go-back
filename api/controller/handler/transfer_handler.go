package handler

import (
	"net/http"
	"simplebank/api/controller/service"
	"simplebank/api/sqlc"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AssignTransferHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := service.NewTransferService(store)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", CreateTransferHandler)
	g.GET("/:id", GetTransferHandler)
	g.GET("/", GetListTransferHandler)

}

func GetTransferHandler(c echo.Context) error {
	service := c.Get("Service").(service.TransferService)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	transfer, err := service.GetTransfer(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transfer)
}

func CreateTransferHandler(c echo.Context) error {
	var err error
	service := c.Get("Service").(service.TransferService)
	params := &sqlc.TransferTxParams{}
	if err = c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(params); err != nil {
		return err
	}
	transfer, err := service.CreateTransfer(params)
	if err != nil {
		//TODO:内部で起きたエラーメッセージをフロントで表示する
		return err
	}
	return c.JSON(http.StatusOK, transfer)
}

func GetListTransferHandler(c echo.Context) error {
	service := c.Get("Service").(service.TransferService)
	params := &sqlc.ListTransfersParams{}
	c.Bind(params)
	transfers, err := service.GetListTransfer(params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transfers)
}
