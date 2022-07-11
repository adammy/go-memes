package font

import (
	"github.com/golang/freetype/truetype"
)

// Repository defines the interface for getting fonts.
type Repository interface {
	// Get a font.
	Get(name string) (*truetype.Font, error)

	// GetPath a path for a font.
	GetPath(name string) (string, error)
}
