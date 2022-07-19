package image

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
)

var _ Getter = (*localGetter)(nil)

type localGetter struct {
	paths map[string]string
}

// NewLocalGetter constructs a localGetter.
func NewLocalGetter(paths map[string]string) *localGetter {
	return &localGetter{
		paths: paths,
	}
}

func (r *localGetter) Get(ID string) (image.Image, error) {
	path, ok := r.paths[ID]
	if !ok {
		return nil, fmt.Errorf("image %s was not found", ID)
	}

	img, err := gg.LoadImage(path)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (r *localGetter) GetPath(ID string) (string, error) {
	if path, ok := r.paths[ID]; ok {
		return path, nil
	}
	return "", fmt.Errorf("image path %s was not found", ID)
}
