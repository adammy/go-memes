package image

import (
	"image"
)

// Uploader defines the interface for an object that uploads an image.
type Uploader interface {
	// UploadPNG pushes an image to some designated storage.
	UploadPNG(filename string, img image.Image) error
}
