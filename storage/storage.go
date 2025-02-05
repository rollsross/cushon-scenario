package storage

import (
	"database/sql"
)

//go:generate mockgen -source=storage/storage.go -destination=storage/mocks/mock_storage.go

type Repository interface {
	CreateAccountAndFund(userId, accountTypeId, fundId string, balance int) error
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
