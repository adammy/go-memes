package repository

const (
	// InMemory denotes to use an in-memory map for templates.
	InMemory Type = "InMemory"

	// Postgres denotes to use a postgres db for templates.
	Postgres Type = "Postgres"
)
