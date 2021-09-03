package api

import "github.com/labstack/echo/v4"




func (server *Server)SetRouter() *echo.Echo{
	e := echo.New()
	//色々middleware設定する用
	return e
}