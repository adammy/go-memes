package meme

// NewRepository constructs a Repository based on the RepositoryType argument.
func NewRepository(t RepositoryType) Repository {
	switch t {
	case InMemoryRepository:
		return NewInMemoryRepository()
	}
	return NewInMemoryRepository()
}
