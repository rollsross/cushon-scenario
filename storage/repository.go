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

func (s *Storage) CreateAccoutAndFund(userId string, data server.CreateAccoutAndFundBody) error {
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
		data.AccountTypeId,
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
		data.Balance,
		data.FundId,
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
