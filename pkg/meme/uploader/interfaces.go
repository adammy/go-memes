package uploader

import (
	"io"
)

// Uploader defines the interface for an object that uploads an image.
type Uploader interface {
	// Upload pushes an image to some designated storage.
	Upload(id string, w io.Writer) error
}
