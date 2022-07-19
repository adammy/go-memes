package template_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/template"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	tests := map[string]struct {
		t template.RepositoryType
	}{
		"inmemory": {
			t: template.InMemoryRepository,
		},
		"postgres": {
			t: template.InMemoryRepository,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := template.NewRepository(tc.t, template.DefaultTemplates)

			assert.NotNil(t, r)
			assert.Implements(t, (*template.Repository)(nil), r)
		})
	}
}
