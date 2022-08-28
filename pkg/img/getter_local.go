package img

import (
	"image"

	"github.com/fogleman/gg"
)

var _ Getter = (*localGetter)(nil)

type localGetter struct{}

// NewLocalGetter constructs a localGetter.
func NewLocalGetter() *localGetter {
	return &localGetter{}
}

func (r *localGetter) Get(path string) (image.Image, error) {
	img, err := gg.LoadImage(path)
	if err != nil {
		return nil, err
	}

	return img, nil
}
