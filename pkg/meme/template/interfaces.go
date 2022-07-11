package template

// Repository defines the interface for template storage.
type Repository interface {
	// Get a meme template from an ID.
	Get(id string) (*Template, error)

	// Create a meme template.
	Create(template *Template) error

	// Delete a meme template from an ID.
	Delete(id string) error
}
