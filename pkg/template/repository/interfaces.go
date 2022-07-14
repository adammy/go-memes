package repository

import (
	"github.com/adammy/memepen-services/pkg/template"
)

// Repository defines the interface for template storage.
type Repository interface {
	// Get a meme template from an ID.
	Get(id string) (*template.Template, error)

	// Create a meme template.
	Create(template *template.Template) error

	// Delete a meme template from an ID.
	Delete(id string) error
}
