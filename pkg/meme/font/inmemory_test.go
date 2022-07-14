package font_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/meme/font"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r := font.NewInMemoryRepository(font.DefaultTestFontPaths)

	assert.NotNil(t, r)
	assert.Implements(t, (*font.Repository)(nil), r)
}

func TestInMemoryRepository_Get(t *testing.T) {
	tests := map[string]struct {
		paths map[string]string
		name  string
		error bool
	}{
		"valid": {
			paths: font.DefaultTestFontPaths,
			name:  "Arial",
		},
		"invalid": {
			paths: font.DefaultTestFontPaths,
			name:  "Fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := font.NewInMemoryRepository(tc.paths)
			textFont, err := r.Get(tc.name)

			if !tc.error {
				assert.NotNil(t, textFont)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestInMemoryRepository_GetPath(t *testing.T) {
	tests := map[string]struct {
		paths        map[string]string
		name         string
		expectedPath string
		error        bool
	}{
		"valid": {
			paths:        font.DefaultTestFontPaths,
			name:         "Arial",
			expectedPath: "../../../assets/fonts/arial.ttf",
		},
		"invalid": {
			paths: font.DefaultTestFontPaths,
			name:  "Fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := font.NewInMemoryRepository(tc.paths)
			path, err := r.GetPath(tc.name)

			if !tc.error {
				assert.Equal(t, tc.expectedPath, path)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
