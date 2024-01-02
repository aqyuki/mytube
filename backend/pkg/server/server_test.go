package server

import (
	"context"
	"testing"

	"github.com/aqyuki/mytube/backend/pkg/account"
	"github.com/aqyuki/mytube/backend/pkg/logging"
	"github.com/aqyuki/mytube/backend/pkg/session"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	m := &Modules{
		AccountService: struct{ account.Service }{},
		SessionManager: &session.Manager{},
	}

	ctx := logging.WithLogger(context.Background(), logging.NewLogger())
	server := New(ctx, m)
	assert.EqualValues(t, m.AccountService, server.accountService, "account service should be set")
	assert.EqualValues(t, m.SessionManager, server.sessionManager, "session manager should be set")
	assert.NotNil(t, server.server, "server should be set")
}
