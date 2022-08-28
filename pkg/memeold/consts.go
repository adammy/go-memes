package memeold

const (
	// InMemoryRepository denotes to use an in-memory map for memes.
	InMemoryRepository RepositoryType = "InMemory"

	// PostgresRepository denotes to use a postgres db for memes.
	PostgresRepository RepositoryType = "Postgres"
)
