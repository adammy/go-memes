package template_test

import (
	"testing"

	"github.com/adammy/go-memes/pkg/meme/template"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r, err := template.NewInMemoryRepository("", nil)

	assert.NotNil(t, r)
	assert.Implements(t, (*template.Repository)(nil), r)
	assert.NoError(t, err)
}

func TestRepository_Get(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid": {
			ID: "yall-got-any-more-of-them",
		},
		"invalid": {
			ID:    "fake",
			error: true,
		},
		"valid with custom": {
			templates: map[string]*template.Template{
				"muh-meme": {},
			},
			ID: "muh-meme",
		},
		"invalid with custom": {
			templates: map[string]*template.Template{},
			ID:        "fake",
			error:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := template.NewInMemoryRepository("", tc.templates)
			tmpl, err := r.Get(tc.ID)

			if !tc.error {
				assert.NotNil(t, tmpl)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestRepository_Create(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid": {
			ID: "muh-meme",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := template.NewInMemoryRepository("", tc.templates)
			err := r.Create(&template.Template{
				ID: tc.ID,
			})

			if !tc.error {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid": {
			templates: map[string]*template.Template{
				"muh-meme": {},
			},
			ID: "muh-meme",
		},
		"invalid": {
			templates: map[string]*template.Template{
				"muh-meme": {},
			},
			ID:    "not-muh-meme",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := template.NewInMemoryRepository("", tc.templates)
			err := r.Delete(tc.ID)

			if !tc.error {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
