package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/aqyuki/mytube/backend/pkg/account"
	"github.com/aqyuki/mytube/backend/pkg/logging"
	mid "github.com/aqyuki/mytube/backend/pkg/middleware"
	ses "github.com/aqyuki/mytube/backend/pkg/session"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server provides the API for the backend
type Server struct {
	// services
	accountService account.Service

	// session store
	sessionManager *ses.Manager

	// server
	server *echo.Echo
}

func New(ctx context.Context, m *Modules, cnf *Config) *Server {
	var s Server

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// set middlewares
	// NOTE: maybe we not required use CORS middlewares
	e.Use(middleware.Recover())

	e.Use(mid.NewLogger(logging.FromContext(ctx)))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(cnf.SessionSecret))))
	e.Use(mid.NewStoreLogger(logging.FromContext(ctx)))

	// set routing
	s.setOAuthRouting(e)

	s.server = e
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
