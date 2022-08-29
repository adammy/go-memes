package memeold_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r := memeold.NewInMemoryRepository()

	assert.NotNil(t, r)
	assert.Implements(t, (*memeold.Repository)(nil), r)
}
