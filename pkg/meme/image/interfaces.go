package image

import (
	"image"
)

// Repository defines the interface for getting images.
type Repository interface {
	// Get an image.
	Get(ID string) (image.Image, error)

	// GetPath gets a path for an image.
	GetPath(ID string) (string, error)
}
