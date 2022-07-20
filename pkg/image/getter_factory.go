package image

import (
	"github.com/rs/zerolog/log"
)

// NewGetter constructs a Getter based on the GetterType argument.
func NewGetter(t GetterType, paths map[string]string) Getter {
	log.Info().Msgf("constructing %s image getter", t)

	switch t {
	case LocalGetter:
		return NewLocalGetter(paths)
	}
	return NewLocalGetter(paths)
}
