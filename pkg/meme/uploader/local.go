package uploader

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

var _ Uploader = (*localUploader)(nil)

type localUploader struct {
	basePath string
}

// NewLocalUploader constructs a localUploader.
func NewLocalUploader(basePath string) (*localUploader, error) {
	return &localUploader{basePath: basePath}, nil
}

func (u *localUploader) UploadPNG(path string, img image.Image) error {
	resolvedPath := filepath.Join(u.basePath, path+".png")
	file, err := os.Create(resolvedPath)
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
