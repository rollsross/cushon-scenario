package storage

import "database/sql"

type Repository interface {
	CreateAccoutAndFund()
	GetAccountAndFund()
}

type Storage struct {
	db *sql.DB
}

// TODO: change to return the interface when methods are implemented
func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}
