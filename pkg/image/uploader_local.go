package image

import (
	"image"
	"image/png"
	"os"
)

var _ Uploader = (*localUploader)(nil)

type localUploader struct{}

// NewLocalUploader constructs a localUploader.
func NewLocalUploader() *localUploader {
	return &localUploader{}
}

func (u *localUploader) UploadPNG(path string, img image.Image) error {
	file, err := os.Create(path + ".png")
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
