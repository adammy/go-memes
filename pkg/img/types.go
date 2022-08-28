package img

import (
	"github.com/adammy/memepen-services/pkg/httpapi"
)

// Config defines configuration for the Server.
type Config struct {
	httpapi.ServerConfig `mapstructure:"server"`

	// RepositoryType defines the repository implementation to use.
	RepositoryType RepositoryType `mapstructure:"repository_type"`

	// UploaderType defines the uploader implementation to use.
	UploaderType UploaderType `mapstructure:"uploader_type"`

	// UploaderBasePath defines the path structure to upload images to.
	UploaderBasePath string `mapstructure:"uploader_base_path"`

	// BaseURL denotes the prefix to the full image URL, e.g., https://cdn.memepen.com/images/.
	BaseURL string `mapstructure:"base_url"`

	// MaxSize denotes the max file size for file uploads.
	MaxSize int64 `mapstructure:"max_size"`

	// MaxWidth denotes the max width of uploader images.
	MaxWidth int `mapstructure:"max_width"`

	// MaxHeight denotes the max height of uploader images.
	MaxHeight int `mapstructure:"max_height"`
}

// GetterType defines the type of Getter.
type GetterType string

// UploaderType defines the type of Uploader.
type UploaderType string

// RepositoryType denotes the type of Repository to use for images.
type RepositoryType string
