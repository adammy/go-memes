package image_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/image"
	"github.com/stretchr/testify/assert"
)

func TestNewUploader(t *testing.T) {
	tests := map[string]struct {
		t image.UploaderType
	}{
		"local": {
			t: image.LocalUploader,
		},
		"noop": {
			t: image.NoopUploader,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			u := image.NewUploader(tc.t)

			assert.NotNil(t, u)
			assert.Implements(t, (*image.Uploader)(nil), u)
		})
	}
}
