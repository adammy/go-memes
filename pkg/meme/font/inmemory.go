package font

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/golang/freetype/truetype"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	basePath string
	paths    map[string]string
	fonts    map[string]*truetype.Font
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository(basePath string, paths map[string]string) (*inMemoryRepository, error) {
	var (
		resolvedPaths map[string]string
	)
	if paths != nil {
		resolvedPaths = paths
	} else {
		resolvedPaths = DefaultFonts
	}

	fonts := map[string]*truetype.Font{}
	for name, path := range resolvedPaths {
		font, _ := getFont(filepath.Join(basePath, path))
		if font != nil {
			fonts[name] = font
		}
	}

	return &inMemoryRepository{
		basePath: basePath,
		paths:    resolvedPaths,
		fonts:    fonts,
	}, nil
}

func (r *inMemoryRepository) Get(name string) (*truetype.Font, error) {
	if font, ok := r.fonts[name]; ok {
		return font, nil
	}
	return nil, fmt.Errorf("font %s was not found", name)
}

func (r *inMemoryRepository) GetPath(name string) (string, error) {
	if path, ok := r.paths[name]; ok {
		return filepath.Join(r.basePath, path), nil
	}
	return "", fmt.Errorf("font path %s was not found", name)
}

func getFont(path string) (*truetype.Font, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return font, err
}
