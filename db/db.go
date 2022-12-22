package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ctxKey int

const (
	dbKey          ctxKey = 0
	defaultTimeout        = 1 * time.Second
)

type Storer interface {
	ListData(brandID string) (data map[string]Verify, err error)
	FindID(ctx context.Context, brandID string, releaseID string) (exists bool, err error)
	ListMonths(releaseID string) (months []string, err error)
	FindIntegrationID(brandID string, integrationID string, size string, sku string, colourcode string) (exists bool, err error)
}

type store struct {
	db *sqlx.DB
}

var db1 *sqlx.DB

func newContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, dbKey, tx)
}

func Transact(ctx context.Context, dbx *sqlx.DB, opts *sql.TxOptions, txFunc func(context.Context) error) (err error) {
	tx, err := dbx.BeginTxx(ctx, opts)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = errors.WithStack(p)
			default:
				err = errors.Errorf("%s", p)
			}
		}
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				err = errors.WithStack(e)
			}
			return
		}
		err = errors.WithStack(tx.Commit())
	}()

	ctxWithTx := newContext(ctx, tx)
	err = WithDefaultTimeout(ctxWithTx, txFunc)
	return err
}

func WithTimeout(ctx context.Context, timeout time.Duration, op func(ctx context.Context) error) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return op(ctxWithTimeout)
}

func WithDefaultTimeout(ctx context.Context, op func(ctx context.Context) error) (err error) {
	return WithTimeout(ctx, defaultTimeout, op)
}

func NewStorer(d *sqlx.DB) Storer {
	return &store{
		db: d,
	}
}
