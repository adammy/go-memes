package img

import (
	"time"
)

const (
	// PNG denotes the png extension.
	PNG string = "png"

	// JPG denotes the jpg extension.
	JPG string = "jpg"

	// JPEG denotes the jpeg extension.
	JPEG string = "jpeg"

	// LocalGetter denotes to use the os filesystem to get images.
	LocalGetter GetterType = "Local"

	// LocalUploader denotes to upload items to the os filesystem.
	LocalUploader UploaderType = "Local"

	// NoopUploader denotes to not upload items and fake success.
	NoopUploader UploaderType = "Noop"

	// InMemoryRepository denotes to use an in-memory map for templates.
	InMemoryRepository RepositoryType = "InMemoryRepository"
)

var (
	// DefaultImages defines the images available for meme templates.
	DefaultImages = map[string]*Image{
		"yall-got-any-more-of-them": {
			ID:        "yall-got-any-more-of-them",
			CreatedOn: time.Now(),
			URL:       "http://localhost:8080/assets/images/yall-got-any-more-of-them.png",
			Width:     600,
			Height:    471,
		},
		"two-buttons": {
			ID:        "two-buttons",
			CreatedOn: time.Now(),
			URL:       "http://localhost:8080/assets/images/two-buttons.png",
			Width:     500,
			Height:    756,
		},
	}
)
