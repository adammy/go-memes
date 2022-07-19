package font

import (
	"github.com/golang/freetype/truetype"
)

// Getter defines the interface for a font getter.
type Getter interface {
	// Get a font.
	Get(name string) (*truetype.Font, error)

	// GetPath for a font.
	GetPath(name string) (string, error)
}
