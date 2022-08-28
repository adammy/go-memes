package img

import (
	"context"
	"image"
	"image/png"
	"os"
)

var _ Uploader = (*localUploader)(nil)

type localUploader struct {
	basePath string
}

// NewLocalUploader constructs a localUploader.
func NewLocalUploader(basePath string) *localUploader {
	return &localUploader{
		basePath: basePath,
	}
}

func (u *localUploader) UploadPNG(_ context.Context, filename string, img image.Image) error {
	file, err := os.Create(u.basePath + filename + "." + PNG)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}
