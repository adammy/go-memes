package repository

import (
	"github.com/adammy/memepen-services/pkg/template"
)

// NewRepository constructs a Repository based on the Type argument.
func NewRepository(t Type, templates map[string]*template.Template) Repository {
	switch t {
	case InMemory:
		return NewInMemoryRepository(templates)
	case Postgres:
		return NewInMemoryRepository(templates)
	}
	return NewInMemoryRepository(templates)
}
