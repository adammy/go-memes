package repository

import (
	"github.com/adammy/memepen-services/pkg/template"
)

func NewRepository(t Type, templates map[string]*template.Template) Repository {
	switch t {
	case InMemory:
		return NewInMemoryRepository(templates)
	case Postgres:
		return NewInMemoryRepository(templates)
	}
	return NewInMemoryRepository(templates)
}
