package meme

import (
	"time"
)

// RepositoryType denotes the type of Repository to use for memes.
type RepositoryType string

// Meme defines a meme.
type Meme struct {
	// ID is a unique identifier for the meme.
	ID string `json:"id"`

	// ImageID define the base image used for the template.
	Image Image `json:"image"`

	// NSFW defines if the meme is inappropriate for work or not.
	NSFW bool `json:"nsfw"`

	// CreatedOn defines the date/time the meme was created.
	CreatedOn time.Time `json:"createdOn"`

	// Text defines the text used in the meme. Stored here for future indexing and searching.
	Text []string `json:"text"`

	// UserID defines the user who created the meme.
	UserID string `json:"userId"`

	// TemplateID defines the template that was used to create this meme.
	TemplateID string `json:"templateId"`
}

type Image struct {
	// Path is the path to the image. This will be a partial path that is missing a base URL (e.g. assets/my-img.png).
	Path string `json:"path"`

	// Width is the width of the image. We want the precalculated and stored to avoid continually calculating this.
	Width uint16 `json:"width"`

	// Height is the height of the image. We want the precalculated and stored to avoid continually calculating this.
	Height uint16 `json:"height"`
}

// CreateMemeFromTemplate informs the Service on how to create a meme.
type CreateMemeFromTemplate struct {
	// Text defines the text used in the meme. Stored here for future indexing and searching.
	Text []string `json:"text" binding:"required"`
}
