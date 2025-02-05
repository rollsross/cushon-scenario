package storage

import (
	"database/sql"

	"github.com/rodionross/cushon-scenario/server"
)

type Repository interface {
	CreateAccoutAndFund(userId string, data server.CreateAccoutAndFundBody) error
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
