package font

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
)

var _ Repository = (*inMemoryRepository)(nil)

type inMemoryRepository struct {
	fontPaths map[string]string
	fonts     map[string]*truetype.Font
}

// NewInMemoryRepository constructs an inMemoryRepository.
func NewInMemoryRepository(fontPaths map[string]string) (*inMemoryRepository, error) {
	var (
		resolvedFontPaths map[string]string
	)
	if fontPaths != nil {
		resolvedFontPaths = fontPaths
	} else {
		resolvedFontPaths = DefaultFonts
	}

	fonts := map[string]*truetype.Font{}
	for name, path := range resolvedFontPaths {
		font, _ := getFont(path)
		if font != nil {
			fonts[name] = font
		}
	}

	return &inMemoryRepository{
		fontPaths: resolvedFontPaths,
		fonts:     fonts,
	}, nil
}

func (r *inMemoryRepository) Get(name string) (*truetype.Font, error) {
	if font, ok := r.fonts[name]; ok {
		return font, nil
	}
	return nil, fmt.Errorf("font %s was not found", name)
}

func (r *inMemoryRepository) GetPath(name string) (string, error) {
	if path, ok := r.fontPaths[name]; ok {
		return path, nil
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
