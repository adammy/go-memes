package meme

// Meme defines a user-created meme.
type Meme struct {
	// Unique identifier for the meme.
	ID string

	// The path to the image. For locally storing, should be "my-image.png". For cloud storing, it'll be some CDN URL.
	ImgPath string

	// The identifier for the template used to create the meme.
	TemplateID string

	// The text used in the meme. Stored here for future indexing and searching.
	Text []string
}
