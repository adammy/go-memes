package font

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
)

var _ Getter = (*localGetter)(nil)

type localGetter struct {
	paths map[string]string
	fonts map[string]*truetype.Font
}

// NewLocalGetter constructs an localGetter.
func NewLocalGetter(paths map[string]string) *localGetter {
	fonts := map[string]*truetype.Font{}
	for name, path := range paths {
		font, _ := getFont(path)
		if font != nil {
			fonts[name] = font
		}
	}

	return &localGetter{
		paths: paths,
		fonts: fonts,
	}
}

func (r *localGetter) Get(name string) (*truetype.Font, error) {
	if font, ok := r.fonts[name]; ok {
		return font, nil
	}
	return nil, ErrNotFound
}

func (r *localGetter) GetPath(name string) (string, error) {
	if path, ok := r.paths[name]; ok {
		return path, nil
	}
	return "", ErrNotFound
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
