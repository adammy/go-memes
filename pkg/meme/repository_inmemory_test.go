package meme_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r := meme.NewInMemoryRepository()

	assert.NotNil(t, r)
	assert.Implements(t, (*meme.Repository)(nil), r)
}
