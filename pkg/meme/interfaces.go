package meme

import (
	"io"
)

// ImageUploader defines the interface for an object that uploads an image.
type ImageUploader interface {
	// Upload pushes an image to some designated storage.
	Upload(id string, w io.Writer) error
}

// MemeRepository defines the interface for an object that stores memes.
type MemeRepository interface {
	// Get a meme from an ID.
	Get(id string) (Meme, error)

	// Create a meme.
	Create(meme *Meme) error

	// Delete a meme from an ID.
	Delete(id string) error
}

// TemplateRepository defines the interface for an object that stores meme templates.
type TemplateRepository interface {
	// Get a meme template from an ID.
	Get(id string) (*Template, error)

	// Create a meme template.
	Create(template *Template) error

	// Delete a meme template from an ID.
	Delete(id string) error
}
