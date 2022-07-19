package template

// NewRepository constructs a Repository based on the RepositoryType argument.
func NewRepository(t RepositoryType, templates map[string]*Template) Repository {
	switch t {
	case InMemoryRepository:
		return NewInMemoryRepository(templates)
	case PostgresRepository:
		return NewInMemoryRepository(templates)
	}
	return NewInMemoryRepository(templates)
}
