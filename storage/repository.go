package storage

import "time"

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
