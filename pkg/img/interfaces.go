package img

import (
	"context"
	"image"
)

// Getter defines the interface for getting an image.Image.
type Getter interface {
	// Get an image.
	Get(path string) (image.Image, error)
}

// Repository defines the interface for Image storage.
type Repository interface {
	// Get an Image from an ID.
	Get(ctx context.Context, ID string) (*Image, error)

	// Create an Image.
	Create(ctx context.Context, template *Image) error

	// Delete an Image from an ID.
	Delete(ctx context.Context, ID string) error
}

// Uploader defines the interface for an object that uploads an image.
type Uploader interface {
	// UploadPNG pushes an image to some designated storage.
	UploadPNG(ctx context.Context, filename string, img image.Image) error
}
