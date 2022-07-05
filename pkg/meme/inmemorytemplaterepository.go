package meme

import (
	"fmt"
)

var _ templateRepository = (*inMemoryTemplateRepository)(nil)

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
		"456": {
			ImgPath: "assets/templates/two-buttons.png",
			Width:   500,
			Height:  756,
			TextStyle: []TextStyle{
				{
					X:     50,
					Y:     100,
					Width: 100,
					Font: Font{
						Family: "Impact",
						Size:   20,
						Color:  "#000000",
					},
					Rotation: &Rotation{
						Degrees: -10,
					},
				},
				{
					X:     240,
					Y:     100,
					Width: 100,
					Font: Font{
						Family: "Impact",
						Size:   20,
						Color:  "#000000",
					},
					Stroke: &Stroke{
						Size:  2,
						Color: "#FF0000",
					},
					Rotation: &Rotation{
						Degrees: -10,
					},
				},
				{
					X:     20,
					Y:     675,
					Width: 460,
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
	templates map[string]Template
}

// NewInMemoryTemplateRepository constructs an inMemoryTemplateRepository.
func NewInMemoryTemplateRepository() (*inMemoryTemplateRepository, error) {
	return &inMemoryTemplateRepository{
		templates: templates,
	}, nil
}

// Get a meme template from an ID.
func (r *inMemoryTemplateRepository) get(id string) (*Template, error) {
	if template, ok := r.templates[id]; ok {
		return &template, nil
	}
	return nil, fmt.Errorf("template %s was not found", id)
}

// Create a meme template.
func (r *inMemoryTemplateRepository) create(template *Template) error {
	r.templates[template.ID] = *template
	return nil
}

// Delete a meme template from an ID.
func (r *inMemoryTemplateRepository) delete(id string) error {
	if _, ok := r.templates[id]; ok {
		delete(r.templates, id)
		return nil
	}
	return fmt.Errorf("template %s was not found", id)
}
