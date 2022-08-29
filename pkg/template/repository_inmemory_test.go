package template_test

import (
	"context"
	"testing"

	"github.com/adammy/memepen-services/pkg/template"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	r := template.NewInMemoryRepository(template.DefaultTemplates)

	assert.NotNil(t, r)
	assert.Implements(t, (*template.Repository)(nil), r)
}

func TestInMemoryRepository_Get(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid": {
			templates: template.DefaultTemplates,
			ID:        "yall-got-any-more-of-them",
		},
		"invalid": {
			templates: template.DefaultTemplates,
			ID:        "fake",
			error:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := template.NewInMemoryRepository(tc.templates)
			tmpl, err := r.Get(context.Background(), tc.ID)

			if !tc.error {
				assert.NotNil(t, tmpl)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestInMemoryRepository_Create(t *testing.T) {
	tests := map[string]struct {
		templates map[string]*template.Template
		ID        string
		error     bool
	}{
		"valid": {
			templates: template.DefaultTemplates,
			ID:        "muh-meme",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := template.NewInMemoryRepository(tc.templates)
			err := r.Create(context.Background(), &template.Template{
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

func TestInMemoryRepository_Delete(t *testing.T) {
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
			r := template.NewInMemoryRepository(tc.templates)
			err := r.Delete(context.Background(), tc.ID)

			if !tc.error {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
