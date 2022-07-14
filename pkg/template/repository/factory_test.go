package repository_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/template"
	"github.com/adammy/memepen-services/pkg/template/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	tests := map[string]struct {
		t repository.Type
	}{
		"inmemory": {
			t: repository.InMemory,
		},
		"postgres": {
			t: repository.InMemory,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := repository.NewRepository(tc.t, template.DefaultTemplates)

			assert.NotNil(t, r)
			assert.Implements(t, (*repository.Repository)(nil), r)
		})
	}
}
