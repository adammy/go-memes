package image

func NewRepository(t RepositoryType, paths map[string]string) Repository {
	switch t {
	case Local:
		return NewLocalRepository(paths)
	}
	return NewLocalRepository(paths)
}
