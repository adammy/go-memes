package meme

import (
	"io"
)

// imageUploader defines the interface for an object that uploads an image.
type imageUploader interface {
	// Upload pushes an image to some designated storage.
	upload(id string, w io.Writer) error
}

// memeRepository defines the interface for an object that stores memes.
type memeRepository interface {
	// get a meme from an ID.
	get(id string) (Meme, error)

	// create a meme.
	create(meme *Meme) error

	// delete a meme from an ID.
	delete(id string) error
}

// templateRepository defines the interface for an object that stores meme templates.
type templateRepository interface {
	// get a meme template from an ID.
	get(id string) (*Template, error)

	// Create a meme template.
	create(template *Template) error

	// Delete a meme template from an ID.
	delete(id string) error
}
