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
		"valid get": {
			ID: "yall-got-any-more-of-them",
		},
		"invalid get": {
			ID:    "fake",
			error: true,
		},
		"valid get with custom": {
			templates: map[string]*template.Template{
				"muh-meme": {},
			},
			ID: "muh-meme",
		},
		"invalid get with custom": {
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

func TestRepository_Create(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid creation": {
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
		"valid deletion": {
			templates: map[string]*template.Template{
				"muh-meme": {},
			},
			ID: "muh-meme",
		},
		"invalid deletion": {
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
