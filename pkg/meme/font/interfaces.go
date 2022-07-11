package font

import (
	"github.com/golang/freetype/truetype"
)

// Repository defines the interface for font storage.
type Repository interface {
	// Get returns a font.
	Get(name string) (*truetype.Font, error)

	// GetPath returns the path for a font.
	GetPath(name string) (string, error)
}
