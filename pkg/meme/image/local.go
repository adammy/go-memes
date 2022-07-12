package image

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
)

var _ Repository = (*localRepository)(nil)

type localRepository struct {
	paths map[string]string
}

// NewLocalRepository constructs a localRepository.
func NewLocalRepository(paths map[string]string) *localRepository {
	return &localRepository{
		paths: paths,
	}
}

func (r *localRepository) Get(ID string) (image.Image, error) {
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

func (r *localRepository) GetPath(ID string) (string, error) {
	if path, ok := r.paths[ID]; ok {
		return path, nil
	}
	return "", fmt.Errorf("image path %s was not found", ID)
}
