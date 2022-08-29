package img

import (
	"context"
	"image"
)

var _ Uploader = (*noopUploader)(nil)

type noopUploader struct{}

// NewNoopUploader constructs a noopUploader.
func NewNoopUploader() *noopUploader {
	return &noopUploader{}
}

func (u *noopUploader) UploadPNG(_ context.Context, _ string, _ image.Image) error {
	return nil
}
