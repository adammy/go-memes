package template

import (
	"context"
	"time"

	"github.com/adammy/memepen-services/pkg/httpapi"
	"github.com/google/uuid"
)

// Service provides the functionality for Template.
type Service struct {
	repository Repository
}

// GetTemplate gets the requested Template.
func (s *Service) GetTemplate(ctx context.Context, templateID string) (*Template, error) {
	template, err := s.repository.Get(ctx, templateID)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateTemplate creates a Template.
func (s *Service) CreateTemplate(ctx context.Context, create CreateTemplate) (*Template, error) {
	if len(create.TextStyles) != len(create.DefaultText) || len(create.TextStyles) == 0 || len(create.DefaultText) == 0 {
		return nil, httpapi.ErrBadRequest
	}

	template := &Template{
		ID:          uuid.NewString(),
		Name:        create.Name,
		CreatedOn:   time.Now(),
		ImageID:     create.ImageID,
		TextStyles:  create.TextStyles,
		DefaultText: create.DefaultText,
	}

	if err := s.repository.Create(ctx, template); err != nil {
		return nil, err
	}

	return template, nil
}

// UpdateTemplate updates the requested Template.
func (s *Service) UpdateTemplate(ctx context.Context, templateID string, create CreateTemplate) (*Template, error) {
	if len(create.TextStyles) != len(create.DefaultText) || len(create.TextStyles) == 0 || len(create.DefaultText) == 0 {
		return nil, httpapi.ErrBadRequest
	}

	template := &Template{
		ID:          templateID,
		Name:        create.Name,
		CreatedOn:   time.Now(),
		ImageID:     create.ImageID,
		TextStyles:  create.TextStyles,
		DefaultText: create.DefaultText,
	}

	if err := s.repository.Update(ctx, templateID, template); err != nil {
		return nil, err
	}

	return template, nil
}

// DeleteTemplate deletes the requested Template.
func (s *Service) DeleteTemplate(ctx context.Context, templateID string) error {
	if err := s.repository.Delete(ctx, templateID); err != nil {
		return err
	}
	return nil
}
