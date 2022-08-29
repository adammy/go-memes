package memeold_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	tests := map[string]struct {
		t memeold.RepositoryType
	}{
		"inmemory": {
			t: memeold.InMemoryRepository,
		},
		"postgres": {
			t: memeold.InMemoryRepository,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := memeold.NewRepository(tc.t)

			assert.NotNil(t, r)
			assert.Implements(t, (*memeold.Repository)(nil), r)
		})
	}
}
