package image_test

import (
	"image"
	"testing"

	img "github.com/adammy/memepen-services/pkg/image"
	"github.com/stretchr/testify/assert"
)

func TestNewNoopUploader(t *testing.T) {
	u := img.NewNoopUploader()

	assert.NotNil(t, u)
	assert.Implements(t, (*img.Uploader)(nil), u)
}

func TestNoopUploader_UploadPNG(t *testing.T) {
	u := img.NewNoopUploader()
	err := u.UploadPNG("", image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 1, Y: 1},
	}))

	assert.NoError(t, err)
}
