package storage

import (
	"database/sql"
)

type Repository interface {
	CreateAccoutAndFund(userId, accountTypeId, fundId string, balance int) error
	GetAccountAndFund(userId string) (*AccountFund, error)
}

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &Storage{
		db: db,
	}
}
