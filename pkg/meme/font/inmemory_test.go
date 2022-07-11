package font_test

import (
	"testing"

	"github.com/adammy/go-memes/pkg/meme/font"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r, _ := font.NewInMemoryRepository("", nil)

	assert.NotNil(t, r)
	assert.Implements(t, (*font.Repository)(nil), r)
}

func TestRepository_Get(t *testing.T) {
	tests := map[string]struct {
		name  string
		error bool
	}{
		"valid font": {
			name: "Arial",
		},
		"invalid font": {
			name:  "Fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := font.NewInMemoryRepository("../../../", nil)
			font, err := r.Get(tc.name)

			if !tc.error {
				assert.NotNil(t, font)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestRepository_GetPath(t *testing.T) {
	tests := map[string]struct {
		fonts        map[string]string
		name         string
		expectedPath string
		error        bool
	}{
		"valid font": {
			name:         "Arial",
			expectedPath: "assets/fonts/arial.ttf",
		},
		"invalid font": {
			name:  "Fake",
			error: true,
		},
		"valid font with custom": {
			fonts: map[string]string{
				"Poppins": "my/custom/font/poppins.ttf",
			},
			name:         "Poppins",
			expectedPath: "my/custom/font/poppins.ttf",
		},
		"invalid font with custom": {
			fonts: map[string]string{
				"Poppins": "my/custom/font/poppins.ttf",
			},
			name:  "Fake",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := font.NewInMemoryRepository("", tc.fonts)
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
