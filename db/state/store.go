package state

import (
	"context"
	"time"
)

type Account struct {
	ID             int32
	DocumentNumber string
	Balance        float64
	CreatedAt      time.Time
}

type OperationType struct {
	ID        string
	CreatedAt time.Time
}

type Transaction struct {
	AccountID       int32
	Amount          float64
	TransactionType string
	EventDate       time.Time
}

type Store interface {
	CreateAccount(ctx context.Context, dn string) (Account, error)
	DeleteAccount(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	UpdateAccount(ctx context.Context, a Account) (Account, error)
	CreateTransaction(ctx context.Context, t Transaction) (Transaction, error)
}
