package image_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/meme/image"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalRepository(t *testing.T) {
	r := image.NewLocalRepository(image.DefaultTestImagePaths)

	assert.NotNil(t, r)
	assert.Implements(t, (*image.Repository)(nil), r)
}

func TestLocalRepository_Get(t *testing.T) {
	tests := map[string]struct {
		paths map[string]string
		ID    string
		error bool
	}{
		"valid": {
			paths: image.DefaultTestImagePaths,
			ID:    "two-buttons",
		},
		"invalid": {
			paths: image.DefaultTestImagePaths,
			ID:    "fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := image.NewLocalRepository(tc.paths)
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

func TestLocalRepository_GetPath(t *testing.T) {
	tests := map[string]struct {
		paths map[string]string
		ID    string
		error bool
	}{
		"valid": {
			paths: image.DefaultTestImagePaths,
			ID:    "two-buttons",
		},
		"invalid": {
			paths: image.DefaultTestImagePaths,
			ID:    "fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := image.NewLocalRepository(tc.paths)
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
