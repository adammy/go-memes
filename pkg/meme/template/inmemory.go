package template

import (
	"fmt"
	"path/filepath"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	basePath  string
	templates map[string]*Template
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository(basePath string, templates map[string]*Template) (*inMemoryRepository, error) {
	var (
		resolvedTemplates map[string]*Template
	)
	if templates != nil {
		resolvedTemplates = templates
	} else {
		resolvedTemplates = DefaultTemplates
	}

	for _, template := range resolvedTemplates {
		template.ImgPath = filepath.Join(basePath, template.ImgPath)
	}

	return &inMemoryRepository{
		basePath:  basePath,
		templates: resolvedTemplates,
	}, nil
}

func (r *inMemoryRepository) Get(id string) (*Template, error) {
	if template, ok := r.templates[id]; ok {
		return template, nil
	}
	return nil, fmt.Errorf("template %s was not found", id)
}

func (r *inMemoryRepository) Create(template *Template) error {
	template.ImgPath = filepath.Join(r.basePath, template.ImgPath)
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
