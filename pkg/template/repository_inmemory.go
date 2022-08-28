package template

import (
	"context"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	templates map[string]*Template
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository(templates map[string]*Template) *inMemoryRepository {
	return &inMemoryRepository{
		templates: templates,
	}
}

func (r *inMemoryRepository) Get(_ context.Context, id string) (*Template, error) {
	if template, ok := r.templates[id]; ok {
		return template, nil
	}
	return nil, ErrNotFound
}

func (r *inMemoryRepository) Create(_ context.Context, template *Template) error {
	r.templates[template.ID] = template
	return nil
}

func (r *inMemoryRepository) Update(_ context.Context, id string, template *Template) error {
	if _, ok := r.templates[id]; !ok {
		return ErrNotFound
	}
	r.templates[id] = template
	return nil
}

func (r *inMemoryRepository) Delete(_ context.Context, id string) error {
	if _, ok := r.templates[id]; ok {
		delete(r.templates, id)
		return nil
	}
	return ErrNotFound
}
