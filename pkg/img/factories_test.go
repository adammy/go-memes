package img_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/img"
	"github.com/stretchr/testify/assert"
)

func TestNewUploader(t *testing.T) {
	tests := map[string]struct {
		t img.UploaderType
	}{
		"local": {
			t: img.LocalUploader,
		},
		"noop": {
			t: img.NoopUploader,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			u := img.NewUploader(tc.t, "")

			assert.NotNil(t, u)
			assert.Implements(t, (*img.Uploader)(nil), u)
		})
	}
}
