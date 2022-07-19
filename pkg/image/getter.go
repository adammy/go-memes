package image

import (
	"image"
)

// Getter defines the interface for an image getter.
type Getter interface {
	// Get an image.
	Get(ID string) (image.Image, error)

	// GetPath for an image.
	GetPath(ID string) (string, error)
}
