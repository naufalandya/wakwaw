package db

import (
	"context"
	"database/sql"
)

// Tx represents a database transaction
type Tx interface {
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
