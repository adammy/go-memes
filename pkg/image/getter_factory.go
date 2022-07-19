package image

// NewGetter constructs a Getter based on the GetterType argument.
func NewGetter(t GetterType, paths map[string]string) Getter {
	switch t {
	case LocalGetter:
		return NewLocalGetter(paths)
	}
	return NewLocalGetter(paths)
}
