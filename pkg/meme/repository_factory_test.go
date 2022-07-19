package meme_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	tests := map[string]struct {
		t meme.RepositoryType
	}{
		"inmemory": {
			t: meme.InMemoryRepository,
		},
		"postgres": {
			t: meme.InMemoryRepository,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := meme.NewRepository(tc.t)

			assert.NotNil(t, r)
			assert.Implements(t, (*meme.Repository)(nil), r)
		})
	}
}
