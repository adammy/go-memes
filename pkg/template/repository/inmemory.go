package repository

import (
	"fmt"

	templatePkg "github.com/adammy/memepen-services/pkg/template"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	templates map[string]*templatePkg.Template
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository(templates map[string]*templatePkg.Template) *inMemoryRepository {
	return &inMemoryRepository{
		templates: templates,
	}
}

func (r *inMemoryRepository) Get(id string) (*templatePkg.Template, error) {
	if template, ok := r.templates[id]; ok {
		return template, nil
	}
	return nil, fmt.Errorf("template %s was not found", id)
}

func (r *inMemoryRepository) Create(template *templatePkg.Template) error {
	r.templates[template.ID] = template
	return nil
}

func (r *inMemoryRepository) Delete(id string) error {
	if _, ok := r.templates[id]; ok {
		delete(r.templates, id)
		return nil
	}
	return fmt.Errorf("template %s was not found", id)
}
