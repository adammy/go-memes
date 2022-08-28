package img

import (
	"context"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	images map[string]*Image
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository(images map[string]*Image) *inMemoryRepository {
	return &inMemoryRepository{
		images: images,
	}
}

func (r *inMemoryRepository) Get(_ context.Context, ID string) (*Image, error) {
	if img, ok := r.images[ID]; ok {
		return img, nil
	}
	return nil, ErrNotFound
}

func (r *inMemoryRepository) Create(_ context.Context, img *Image) error {
	r.images[img.ID] = img
	return nil
}

func (r *inMemoryRepository) Delete(_ context.Context, ID string) error {
	if _, ok := r.images[ID]; ok {
		delete(r.images, ID)
		return nil
	}
	return ErrNotFound
}
