package api

import (
	"log"
	"simplebank/api/controller/handler"
	"simplebank/api/middleware"
	"simplebank/api/util"

	"github.com/labstack/echo/v4"
)


func (server *Server) SetRouter() *echo.Echo {
	e := echo.New()
	validator, err := util.NewValidator()
	if err != nil {
		log.Fatal("バリデージョンの設定に失敗しました")
	}
	e.Validator = validator
	middleware.CORS(e)
	{
		g := e.Group("/api/v1", middleware.SetStore(server.store))
		handler.AssignAccountHandler(g.Group("/account"))
	}
	return e
}
