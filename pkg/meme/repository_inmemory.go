package meme

import (
	"context"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	memes map[string]Meme
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		memes: map[string]Meme{},
	}
}

func (r *inMemoryRepository) Get(ctx context.Context, ID string) (*Meme, error) {
	if meme, ok := r.memes[ID]; ok {
		return &meme, nil
	}
	return nil, ErrNotFound
}

func (r *inMemoryRepository) Create(ctx context.Context, meme Meme) error {
	r.memes[meme.ID] = meme
	return nil
}

func (r *inMemoryRepository) Delete(ctx context.Context, ID string) error {
	if _, ok := r.memes[ID]; ok {
		delete(r.memes, ID)
		return nil
	}
	return ErrNotFound
}
