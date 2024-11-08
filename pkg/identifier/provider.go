package identifier

import "github.com/rs/xid"

// Provider is an interface to provide features related to identifier generation.
type Provider interface {
	// Generate generates a new identifier.
	Generate() string
}

// XidIdentifierProvider must be implemented by IdentifierProvider.
var _ Provider = (*XidIdentifierProvider)(nil)

// XidIdentifierProvider is a implementation of IdentifierProvider.
type XidIdentifierProvider struct{}

// NewXidIdentifierProvider creates a new XidIdentifierProvider.
func NewXidIdentifierProvider() *XidIdentifierProvider { return &XidIdentifierProvider{} }

// Generate generates a new identifier.
func (p *XidIdentifierProvider) Generate() string {
	return xid.New().String()
}
