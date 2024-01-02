package server

import (
	"testing"

	"github.com/aqyuki/mytube/backend/pkg/account"
	"github.com/aqyuki/mytube/backend/pkg/session"
	"github.com/stretchr/testify/assert"
)

func Test_Module(t *testing.T) {
	t.Parallel()

	accountService := struct {
		account.Service
	}{}
	s := session.Manager{}

	m := Modules{
		AccountService: &accountService,
		SessionManager: &s,
	}

	server := new(Server)
	m.RegisterServices(server)

	assert.EqualValues(t, &accountService, server.accountService)
	assert.EqualValues(t, &s, server.sessionManager)
}
