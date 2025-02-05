package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/rodionross/cushon-scenario/server"
)

type User struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Account struct {
	Id            string
	AccountTypeId string
	UserId        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type AccountType struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Fund struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountFund struct {
	AccountName string `json:"accountName" example:"Cushon ISA"`
	FundName    string `json:"fundName" example:"Cushon Equities Fund"`
	Balance     int    `json:"balance" example:"2500000"`
}

func (s *Storage) CreateAccoutAndFund(userId, accountTypeId, fundId string, balance int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	accountId, _ := uuid.NewV7()

	insertAccount := `
	INSERT INTO accounts(id, account_types_id, users_id, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?);`
	_, err = tx.Exec(
		insertAccount,
		accountId,
		accountTypeId,
		userId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	insertAccountFund := `
	INSERT INTO accounts_funds(balance, funds_id, accounts_id, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?);`
	_, err = tx.Exec(
		insertAccountFund,
		balance,
		fundId,
		accountId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
