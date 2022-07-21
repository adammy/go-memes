package image

import (
	"github.com/rs/zerolog/log"
)

// NewUploader constructs an UploaderType based on the Type argument.
func NewUploader(t UploaderType) Uploader {
	log.Info().Msgf("constructing %s image uploader", t)

	switch t {
	case LocalUploader:
		return NewLocalUploader()
	case NoopUploader:
		return NewNoopUploader()
	}
	return NewLocalUploader()
}
