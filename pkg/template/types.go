package template

import (
	"github.com/adammy/memepen-services/pkg/httpapi"
)

// Config defines configuration for the Server.
type Config struct {
	httpapi.ServerConfig `mapstructure:"server"`

	// RepositoryType defines the store for templates.
	RepositoryType RepositoryType `mapstructure:"repository_type"`
}

// RepositoryType denotes the type of Repository to use for templates.
type RepositoryType string
