package template_test

import (
	"testing"

	"github.com/adammy/go-memes/pkg/meme/template"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r, _ := template.NewInMemoryRepository("", nil)

	assert.NotNil(t, r)
	assert.Implements(t, (*template.Repository)(nil), r)
}

func TestRepository_Get(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid template": {
			ID: "yall-got-any-more-of-them",
		},
		"invalid template": {
			ID:    "fake",
			error: true,
		},
		"valid template with custom": {
			templates: map[string]*template.Template{
				"muh-meme": {},
			},
			ID: "muh-meme",
		},
		"invalid template with custom": {
			templates: map[string]*template.Template{},
			ID:        "fake",
			error:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := template.NewInMemoryRepository("", tc.templates)
			template, err := r.Get(tc.ID)

			if !tc.error {
				assert.NotNil(t, template)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
