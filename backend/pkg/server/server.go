package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/aqyuki/mytube/backend/pkg/account"
	"github.com/aqyuki/mytube/backend/pkg/session"
	"github.com/labstack/echo/v4"
)

// Server provides the API for the backend
type Server struct {
	// services
	accountService account.Service

	// session store
	sessionManager *session.Manager

	// server
	server *echo.Echo
}

func New(m *Modules) *Server {
	var s Server
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	s.server = server
	m.RegisterServices(&s)
	return &s
}

func (s *Server) Start(port int) error {
	err := s.server.Start(fmt.Sprintf(":%d", port))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
