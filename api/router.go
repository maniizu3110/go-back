package api

import (
	"simplebank/api/controller/handler"
	"simplebank/api/middleware"

	"github.com/labstack/echo/v4"
)

func (server *Server) SetRouter() *echo.Echo {
	e := echo.New()
	middleware.CORS(e)
	{
		g := e.Group("/api/v1", middleware.SetStore(server.store))
		handler.AssignAccountHandler(g.Group("/account"))
	}
	return e
}
