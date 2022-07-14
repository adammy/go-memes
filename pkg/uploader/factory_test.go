package uploader_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/uploader"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	tests := map[string]struct {
		t uploader.Type
	}{
		"local": {
			t: uploader.Local,
		},
		"noop": {
			t: uploader.Noop,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := uploader.NewUploader(tc.t)

			assert.NotNil(t, r)
			assert.Implements(t, (*uploader.Uploader)(nil), r)
		})
	}
}
