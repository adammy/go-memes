package image_test

import (
	"testing"

	"github.com/adammy/go-memes/pkg/meme/image"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalRepository(t *testing.T) {
	r, err := image.NewLocalRepository("", nil)

	assert.NotNil(t, r)
	assert.Implements(t, (*image.Repository)(nil), r)
	assert.NoError(t, err)
}

func TestRepository_Get(t *testing.T) {
	tests := map[string]struct {
		ID    string
		error bool
	}{
		"valid": {
			ID: "two-buttons",
		},
		"invalid": {
			ID:    "fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := image.NewLocalRepository("../../../", nil)
			img, err := r.Get(tc.ID)

			if !tc.error {
				assert.NotNil(t, img)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestRepository_GetPath(t *testing.T) {
	tests := map[string]struct {
		ID    string
		error bool
	}{
		"valid": {
			ID: "two-buttons",
		},
		"invalid": {
			ID:    "fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := image.NewLocalRepository("", nil)
			path, err := r.GetPath(tc.ID)

			if !tc.error {
				assert.NotNil(t, path)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
