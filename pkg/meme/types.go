package meme

import (
	"time"
)

type RepositoryType string

const (
	InMemory RepositoryType = "InMemory"
)

// Meme defines a meme.
type Meme struct {
	// ID is a unique identifier for the meme.
	ID string `json:"id"`

	// TemplateID defines the template that was used to create this meme.
	TemplateID string `json:"templateId"`

	// ImgPath is the path to the image. This will be a partial path that is missing a base URL (e.g. assets/my-img.png).
	ImgPath string `json:"path"`

	// Text defines the text used in the meme. Stored here for future indexing and searching.
	Text []string `json:"text"`

	// UserID defines the user who created the meme.
	UserID string `json:"userId"`

	// CreatedOn defines the date/time the meme was created.
	CreatedOn time.Time `json:"createdOn"`
}

// CreateMemeFromTemplate informs the Service on how to create a meme.
type CreateMemeFromTemplate struct {
	// Text defines the text used in the meme. Stored here for future indexing and searching.
	Text []string `json:"text" binding:"required"`
}
