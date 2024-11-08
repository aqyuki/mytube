package identifier

import (
	"testing"
)

func TestNewXidIdentifierProvider(t *testing.T) {
	t.Parallel()

	provider := NewXidIdentifierProvider()
	if provider == nil {
		t.Error("provider expected to be not nil but received nil")
	}
}

func TestXidIdentifierProvider_Generate(t *testing.T) {
	t.Parallel()

	provider := NewXidIdentifierProvider()
	identifier := provider.Generate()
	if identifier == "" {
		t.Error("identifier expected to be not empty but received empty")
	}
}
