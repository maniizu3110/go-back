package api

import (
	"fmt"
	"simplebank/api/sqlc"
	"simplebank/api/util"
	"simplebank/token"

	"github.com/labstack/echo/v4"
)

type Server struct {
	config     util.Config
	store      sqlc.Store
	tokenMaker token.Maker
	router     *echo.Echo
}

func NewServer(config util.Config, store sqlc.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router = server.SetRouter()
	return server, nil

}
func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
