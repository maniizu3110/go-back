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
	e.Use(middleware.LoggingMiddleware)
	validator, err := util.NewValidator()
	if err != nil {
		log.Fatal("バリデージョンの設定に失敗しました")
	}
	e.Validator = validator
	middleware.CORS(e)
	//TODO:認証が必要なAPIとそれ以外を分ける（authentication）
	{
		//認証不要
		g := e.Group("/api/v1", middleware.SetStore(server.store))
		handler.AssignLoginHandler(g.Group("/login"))
		
	}
	{
		//要認証認証
		g := e.Group("/api/v1", middleware.SetStore(server.store))
		handler.AssignAccountHandler(g.Group("/account"))
		handler.AssignTransferHandler(g.Group("/transfer"))
		handler.AssignUserHandler(g.Group("/user"))
	}
	return e
}
