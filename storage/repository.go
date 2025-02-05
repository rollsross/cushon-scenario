package storage

import (
	"time"

	"github.com/google/uuid"
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

func (s *Storage) GetAccountAndFund(userId string) (*AccountFund, error) {
	getAccountAndFund := `
	SELECT account_types.name, funds.name, accounts_funds.balance
    FROM accounts
    JOIN account_types ON accounts.account_types_id = account_types.id
    JOIN accounts_funds ON accounts.id = accounts_funds.accounts_id
    JOIN funds ON accounts_funds.funds_id = funds.id
    WHERE accounts.users_id = '00a79964-34c2-48ab-88ab-de65427cb960'
    LIMIT 1;`
	rows, err := s.db.Query(getAccountAndFund, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res AccountFund

	for rows.Next() {
		err := rows.Scan(&res.AccountName, &res.FundName, &res.Balance)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}
