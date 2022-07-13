package font

func NewRepository(t RepositoryType, paths map[string]string) Repository {
	switch t {
	case InMemory:
		return NewInMemoryRepository(paths)
	}
	return NewInMemoryRepository(paths)
}
