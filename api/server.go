package api

import (
	db "simplebank/db/sqlc"

	"github.com/labstack/echo/v4"
)

type Server struct {
	store  *db.Store
	router *echo.Echo
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	server.router = server.SetRouter()
	return server

}
func (server *Server) Start(address string){
	server.router.Logger.Fatal(server.router.Start(address))
}
