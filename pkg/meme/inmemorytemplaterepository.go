package meme

import (
	"fmt"
)

var _ TemplateRepository = (*inMemoryTemplateRepository)(nil)

var (
	templates = map[string]Template{
		"123": {
			ImgPath: "assets/templates/yall-got-any-more-of-that.png",
			Width:   600,
			Height:  471,
			TextStyle: []TextStyle{
				{
					X:     10,
					Y:     10,
					Width: 580,
					Font: Font{
						Family: "Impact",
						Size:   40,
						Color:  "#FFFFFF",
					},
					Stroke: &Stroke{
						Size:  4,
						Color: "#000000",
					},
				},
				{
					X:     10,
					Y:     421,
					Width: 580,
					Font: Font{
						Family: "Impact",
						Size:   40,
						Color:  "#FFFFFF",
					},
					Stroke: &Stroke{
						Size:  4,
						Color: "#000000",
					},
				},
			},
		},
	}
)

type inMemoryTemplateRepository struct {
	Templates map[string]Template
}

func NewInMemoryTemplateRepository() (*inMemoryTemplateRepository, error) {
	return &inMemoryTemplateRepository{
		Templates: templates,
	}, nil
}

func (r *inMemoryTemplateRepository) Get(id string) (*Template, error) {
	if template, ok := r.Templates[id]; ok {
		return &template, nil
	}
	return nil, fmt.Errorf("template %s was not found", id)
}

func (r *inMemoryTemplateRepository) Create(template *Template) error {
	r.Templates[template.ID] = *template
	return nil
}

func (r *inMemoryTemplateRepository) Delete(id string) error {
	if _, ok := r.Templates[id]; ok {
		delete(r.Templates, id)
		return nil
	}
	return fmt.Errorf("template %s was not found", id)
}
