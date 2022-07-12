package uploader

import (
	"image"
)

var _ Uploader = (*noopUploader)(nil)

type noopUploader struct{}

// NewNoopUploader constructs a noopUploader.
func NewNoopUploader() *noopUploader {
	return &noopUploader{}
}

func (u *noopUploader) UploadPNG(_ string, _ image.Image) error {
	return nil
}
