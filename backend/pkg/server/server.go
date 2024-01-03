package server

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrServerNotInitialized = errors.New("server is not initialized")
	ErrServerConfigNotInit  = errors.New("server config is not initialized")
)

// Server is the server.
type Server struct {
	server *echo.Echo
	Config *Config
}

func New(cnf *Config) *Server {
	e := echo.New()

	// TOOD: register routes and middleware

	return &Server{
		server: e,
		Config: cnf,
	}
}

// Start starts the server.
func (s *Server) Start() error {
	if s == nil {
		return ErrServerNotInitialized
	} else if s.Config == nil {
		return ErrServerConfigNotInit
	}

	if s.Config.UseTLS {
		return s.startTLS()
	}
	return s.startHttp()
}

func (s *Server) startTLS() error {
	return s.server.Start(s.Config.Addr())
}

func (s *Server) startHttp() error {
	return s.server.StartTLS(s.Config.Addr(), s.Config.CrtPath, s.Config.KeyPath)
}
