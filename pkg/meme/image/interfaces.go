package image

import (
	"image"
)

// Repository defines the interface for image storage.
type Repository interface {
	// Get returns an image.
	Get(ID string) (image.Image, error)

	// GetPath returns the path for an image.
	GetPath(ID string) (string, error)
}
