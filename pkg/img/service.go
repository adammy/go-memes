package img

import (
	"context"
	"errors"
	"image"
	"time"

	"github.com/google/uuid"
)

// Service provides functionality for creating an Image.
type Service struct {
	repository Repository
	uploader   Uploader
	baseURL    string
	maxWidth   int
	maxHeight  int
}

// Get an Image.
func (s *Service) Get(ctx context.Context, imageID string) (*Image, error) {
	img, err := s.repository.Get(ctx, imageID)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// CreateFromUpload creates an Image from an image.Image.
func (s *Service) CreateFromUpload(ctx context.Context, img image.Image) (*Image, error) {
	id := uuid.NewString()
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width > s.maxWidth || height > s.maxHeight {
		return nil, ErrImgSizeTooLarge
	}

	newImg := &Image{
		ID:        id,
		CreatedOn: time.Now(),
		URL:       s.baseURL + id + "." + PNG,
		Width:     width,
		Height:    height,
	}

	if err := s.uploader.UploadPNG(ctx, id, img); err != nil {
		return nil, err
	}

	if err := s.repository.Create(ctx, newImg); err != nil {
		return nil, err
	}

	return newImg, nil
}

// CreateFromRemote creates an Image from a remote URL.
func (s *Service) CreateFromRemote(ctx context.Context, URL string) (*Image, error) {
	return nil, errors.New("not implemented")
}

// Delete an Image.
func (s *Service) Delete(ctx context.Context, imageID string) error {
	if err := s.repository.Delete(ctx, imageID); err != nil {
		return err
	}
	return nil
}
