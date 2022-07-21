package font

import (
	"github.com/rs/zerolog/log"
)

// NewGetter constructs a Getter based on the GetterType argument.
func NewGetter(t GetterType, paths map[string]string) Getter {
	log.Info().Msgf("constructing %s font getter", t)

	switch t {
	case LocalGetter:
		return NewLocalGetter(paths)
	}
	return NewLocalGetter(paths)
}
