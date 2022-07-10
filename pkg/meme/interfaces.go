package meme

// Repository defines the interface for an object that stores memes.
type Repository interface {
	// get a meme from an ID.
	Get(id string) (Meme, error)

	// create a meme.
	Create(meme *Meme) error

	// delete a meme from an ID.
	Delete(id string) error
}
