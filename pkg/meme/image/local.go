package image

import (
	"fmt"
	"image"
	"path/filepath"

	"github.com/fogleman/gg"
)

var _ Repository = (*localRepository)(nil)

type localRepository struct {
	basePath string
	paths    map[string]string
}

// NewLocalRepository constructs a localRepository.
func NewLocalRepository(basePath string, paths map[string]string) (*localRepository, error) {
	var (
		resolvedPaths map[string]string
	)
	if paths != nil {
		resolvedPaths = paths
	} else {
		resolvedPaths = DefaultImages
	}

	return &localRepository{
		basePath: basePath,
		paths:    resolvedPaths,
	}, nil
}

func (r *localRepository) Get(ID string) (image.Image, error) {
	path, ok := r.paths[ID]
	if !ok {
		return nil, fmt.Errorf("image %s was not found", ID)
	}

	img, err := gg.LoadImage(filepath.Join(r.basePath, path))
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (r *localRepository) GetPath(ID string) (string, error) {
	if path, ok := r.paths[ID]; ok {
		return filepath.Join(r.basePath, path), nil
	}
	return "", fmt.Errorf("image path %s was not found", ID)
}
