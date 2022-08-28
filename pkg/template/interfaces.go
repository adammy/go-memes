package template

import (
	"context"
)

// Repository defines the interface for template storage.
type Repository interface {
	// Get a template from an ID.
	Get(ctx context.Context, id string) (*Template, error)

	// Create a template.
	Create(ctx context.Context, template *Template) error

	// Update a template in full.
	Update(ctx context.Context, id string, template *Template) error

	// Delete a template from an ID.
	Delete(ctx context.Context, id string) error
}
