package memeold

// Repository defines the interface for an object that stores memes.
type Repository interface {
	// Get a meme from an ID.
	Get(id string) (*Meme, error)

	// Create a meme.
	Create(meme *Meme) error

	// Delete a meme from an ID.
	Delete(id string) error
}
