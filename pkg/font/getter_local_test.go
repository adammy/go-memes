package font_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/font"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalGetter(t *testing.T) {
	g := font.NewLocalGetter(font.DefaultTestFontPaths)

	assert.NotNil(t, g)
	assert.Implements(t, (*font.Getter)(nil), g)
}

func TestLocalGetter_Get(t *testing.T) {
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
			g := font.NewLocalGetter(tc.paths)
			textFont, err := g.Get(tc.name)

			if !tc.error {
				assert.NotNil(t, textFont)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLocalGetter_GetPath(t *testing.T) {
	tests := map[string]struct {
		paths        map[string]string
		name         string
		expectedPath string
		error        bool
	}{
		"valid": {
			paths:        font.DefaultTestFontPaths,
			name:         "Arial",
			expectedPath: "../../assets/fonts/arial.ttf",
		},
		"invalid": {
			paths: font.DefaultTestFontPaths,
			name:  "Fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			g := font.NewLocalGetter(tc.paths)
			path, err := g.GetPath(tc.name)

			if !tc.error {
				assert.Equal(t, tc.expectedPath, path)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
