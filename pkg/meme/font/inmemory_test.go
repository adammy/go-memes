package font_test

import (
	"testing"

	"github.com/adammy/go-memes/pkg/meme/font"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r, _ := font.NewInMemoryRepository(nil)

	assert.NotNil(t, r)
	assert.Implements(t, (*font.Repository)(nil), r)
}

func TestRepository_GetPath(t *testing.T) {
	tests := map[string]struct {
		fonts    map[string]string
		fontName string
		fontPath string
		error    bool
	}{
		"valid font": {
			fontName: "Arial",
			fontPath: "assets/fonts/arial.ttf",
		},
		"invalid font": {
			fontName: "Fake",
			error:    true,
		},
		"valid font with custom": {
			fonts: map[string]string{
				"Poppins": "my/custom/font/poppins.ttf",
			},
			fontName: "Poppins",
			fontPath: "my/custom/font/poppins.ttf",
		},
		"invalid font with custom": {
			fonts: map[string]string{
				"Poppins": "my/custom/font/poppins.ttf",
			},
			fontName: "Fake",
			error:    true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := font.NewInMemoryRepository(tc.fonts)
			path, err := r.GetPath(tc.fontName)

			if !tc.error {
				assert.Equal(t, tc.fontPath, path)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
