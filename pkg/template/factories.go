package template

// NewRepository constructs a Repository based on the RepositoryType argument.
func NewRepository(t RepositoryType) Repository {
	switch t {
	case InMemoryRepository:
		return NewInMemoryRepository(DefaultTemplates)
	case PostgresRepository:
		return NewInMemoryRepository(DefaultTemplates)
	}
	return NewInMemoryRepository(DefaultTemplates)
}

// NewService constructs a Service.
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}
