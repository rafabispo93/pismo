package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // import driver
	"github.com/pkg/errors"

	"pismo/db/state"
)

// Store provides data interaction methods over a Postgres DB.
type Store struct {
	db *sql.DB

	newID func() uuid.UUID
}

// NewStore instantiates a new Store.
func NewStore(cxn string) (*Store, error) {
	db, err := sql.Open("postgres", cxn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open DB")
	}
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping DB")
	}

	return &Store{
		db:    db,
		newID: uuid.New,
	}, nil
}

// Ping pings the DB.
func (s *Store) Ping(ctx context.Context) error {
	return errors.Wrap(s.db.PingContext(ctx), "failed to ping DB")
}

func (s *Store) CreateAccount(ctx context.Context, dn string) (state.Account, error) {
	q := New(s.db)
	acc, err := q.CreateAccount(ctx, dn)
	if err != nil {
		return state.Account{}, errors.Wrap(err, "failed to create account")
	}
	sa := state.Account{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
		Balance:        acc.Balance,
		CreatedAt:      acc.CreatedAt,
	}
	return sa, nil
}

func (s *Store) DeleteAccount(ctx context.Context, id int32) error {
	q := New(s.db)
	err := q.DeleteAccount(ctx, id)
	return err
}

func (s *Store) GetAccount(ctx context.Context, id int32) (state.Account, error) {
	q := New(s.db)
	acc, err := q.GetAccount(ctx, id)
	if err != nil {
		return state.Account{}, errors.Wrap(err, "failed to get account")
	}

	sa := state.Account{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
		Balance:        acc.Balance,
		CreatedAt:      acc.CreatedAt,
	}
	return sa, nil
}

func (s *Store) UpdateAccount(ctx context.Context, a UpdateAccountParams) (state.Account, error) {
	q := New(s.db)
	acc, err := q.UpdateAccount(ctx, a)
	if err != nil {
		return state.Account{}, errors.Wrap(err, "failed to update account")
	}

	sa := state.Account{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
		Balance:        acc.Balance,
		CreatedAt:      acc.CreatedAt,
	}
	return sa, nil
}

func (s *Store) CreateTransaction(ctx context.Context, ctp CreateTransactionParams) (state.Transaction, error) {
	q := New(s.db)
	t, err := q.CreateTransaction(ctx, ctp)
	if err != nil {
		return state.Transaction{}, errors.Wrap(err, "failed to create transaction")
	}

	st := state.Transaction{
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		EventDate:       t.EventDate,
	}

	return st, nil
}
