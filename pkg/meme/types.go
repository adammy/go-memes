package meme

// Meme defines a meme.
type Meme struct {
	// ID is a unique identifier for the meme.
	ID string

	// ImgPath is the path to the image.
	ImgPath string

	// TemplateID defines the template that was used to create this meme.
	TemplateID string

	// Text defines the text used in the meme. Stored here for future indexing and searching.
	Text []string
}

// CreateMemeFromTemplate informs the Service on how to create a meme.
type CreateMemeFromTemplate struct {
	// Text defines the text used in the meme. Stored here for future indexing and searching.
	Text []string `json:"text" binding:"required"`
}
