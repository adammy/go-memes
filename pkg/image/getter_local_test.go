package image_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/image"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalGetter(t *testing.T) {
	r := image.NewLocalGetter(image.DefaultTestImagePaths)

	assert.NotNil(t, r)
	assert.Implements(t, (*image.Getter)(nil), r)
}

func TestLocalGetter_Get(t *testing.T) {
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
			g := image.NewLocalGetter(tc.paths)
			img, err := g.Get(tc.ID)

			if !tc.error {
				assert.NotNil(t, img)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLocalGetter_GetPath(t *testing.T) {
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
			g := image.NewLocalGetter(tc.paths)
			path, err := g.GetPath(tc.ID)

			if !tc.error {
				assert.NotNil(t, path)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
