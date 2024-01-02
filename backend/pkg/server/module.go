package server

import (
	"github.com/aqyuki/mytube/backend/pkg/account"
	"github.com/aqyuki/mytube/backend/pkg/session"
)

type Modules struct {
	AccountService account.Service
	SessionManager *session.Manager
}

func (m *Modules) RegisterServices(s *Server) {
	s.accountService = m.AccountService
	s.sessionManager = m.SessionManager
}
