package meme

func NewRepository(t RepositoryType) Repository {
	switch t {
	case InMemory:
		return NewInMemoryRepository()
	}
	return NewInMemoryRepository()
}
