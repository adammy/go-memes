package meme

import (
	"context"
)

// Repository defines the interface for an object that stores memes.
type Repository interface {
	// Get a meme from an ID.
	Get(ctx context.Context, ID string) (*Meme, error)

	// Create a meme.
	Create(ctx context.Context, meme Meme) error

	// Delete a meme from an ID.
	Delete(ctx context.Context, ID string) error
}
