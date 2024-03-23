package server

import (
	"github.com/labstack/echo/v4"
	"github.com/vnkdj5/devtools/pkg/config"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	return &Server{
		Echo:   echo.New(),
		Config: cfg,
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
