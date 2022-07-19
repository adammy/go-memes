package meme

import (
	"fmt"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	memes map[string]*Meme
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		memes: map[string]*Meme{},
	}
}

func (r *inMemoryRepository) Get(id string) (*Meme, error) {
	if meme, ok := r.memes[id]; ok {
		return meme, nil
	}
	return nil, fmt.Errorf("meme %s was not found", id)
}

func (r *inMemoryRepository) Create(meme *Meme) error {
	r.memes[meme.ID] = meme
	return nil
}

func (r *inMemoryRepository) Delete(id string) error {
	if _, ok := r.memes[id]; ok {
		delete(r.memes, id)
		return nil
	}
	return fmt.Errorf("meme %s was not found", id)
}
