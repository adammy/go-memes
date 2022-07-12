package uploader_test

import (
	"github.com/adammy/go-memes/pkg/meme/uploader"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestNewNoopUploader(t *testing.T) {
	u := uploader.NewNoopUploader()

	assert.NotNil(t, u)
	assert.Implements(t, (*uploader.Uploader)(nil), u)
}

func TestNoopUploader_UploadPNG(t *testing.T) {
	u := uploader.NewNoopUploader()
	err := u.UploadPNG("", image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 1, Y: 1},
	}))

	assert.NoError(t, err)
}
