// Package database implements methods to interact with the database.
package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

// New creates a new database connection pool using the passed parameters.
func New(dbUser string, dbPassword string, dbHost string, dbPort int, dbName string) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s:%d/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	dbPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s", dsn))
	if err != nil {
		return nil, err
	}

	// test connection
	err = dbPool.Ping(context.Background())
	if err != nil {
		dbPool.Close()

		return nil, err
	}

	return &DB{dbPool}, nil
}

// WithTx runs fn inside a transaction. It commits on success and rolls back on
// any error returned by fn or by Commit itself.
func (db *DB) WithTx(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()
	if err = fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
