package template

func NewRepository(t RepositoryType, templates map[string]*Template) Repository {
	switch t {
	case InMemory:
		return NewInMemoryRepository(templates)
	}
	return NewInMemoryRepository(templates)
}
