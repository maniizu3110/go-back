package api

import (
	"simplebank/api/controller/handler"
	"simplebank/api/middleware"
	"simplebank/api/util"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (server *Server) SetRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggingMiddleware)
	validator, err := util.NewValidator()
	if err != nil {
		logrus.Fatal("バリデージョンの設定に失敗しました")
	}
	if server.config.Env != "prod" {
		e.Debug = true
	}
	e.Validator = validator
	middleware.CORS(e)
	//TODO:認証が必要なAPIとそれ以外を分ける（authentication）
	{
		//認証不要
		g := e.Group("/api/v1",
			middleware.SetStore(server.store),
			middleware.SetConfig(server.config),
			middleware.SetTokenMaker(server.tokenMaker),
		)
		handler.AssignLoginHandler(g.Group("/login"))
		handler.AssignSignHandler(g.Group("/sign"))

	}
	{
		//要認証認証
		g := e.Group("/api/v1", middleware.SetStore(server.store), middleware.AuthMiddleware(server.tokenMaker))
		handler.AssignAccountHandler(g.Group("/account"))
		handler.AssignTransferHandler(g.Group("/transfer"))
		handler.AssignUserHandler(g.Group("/user"))
	}
	return e
}
