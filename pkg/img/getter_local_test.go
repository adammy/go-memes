package img_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/img"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalGetter(t *testing.T) {
	r := img.NewLocalGetter()

	assert.NotNil(t, r)
	assert.Implements(t, (*img.Getter)(nil), r)
}

func TestLocalGetter_Get(t *testing.T) {
	tests := map[string]struct {
		path  string
		error bool
	}{
		"valid": {
			path: "../../assets/images/two-buttons.png",
		},
		"invalid": {
			path:  "fake-path/image.png",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			g := img.NewLocalGetter()
			img, err := g.Get(tc.path)

			if !tc.error {
				assert.NotNil(t, img)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
