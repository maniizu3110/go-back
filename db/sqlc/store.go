package db

import "database/sql"

//Store provides all functions to execure db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{
		Queries: New(db),
		db:db,
	}
}