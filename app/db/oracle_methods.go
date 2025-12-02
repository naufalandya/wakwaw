package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/godror/godror"
)

// ==================== METHODS ====================

func (c *OracleClient) Ping(ctx context.Context) error {
	return c.DB.PingContext(ctx)
}

func (c *OracleClient) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return c.DB.ExecContext(ctx, query, args...)
}

func (c *OracleClient) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return c.DB.QueryContext(ctx, query, args...)
}

func (c *OracleClient) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.DB.QueryRowContext(ctx, query, args...)
}

func (c *OracleClient) Close() error {
	log.Println("[DB] Closing Oracle connection...")
	return c.DB.Close()
}

func (c *OracleClient) Begin(ctx context.Context) (Tx, error) {
	tx, err := c.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &oracleTx{tx: tx}, nil
}

func (t *oracleTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return t.tx.ExecContext(ctx, query, args...)
}

func (t *oracleTx) Commit(ctx context.Context) error {
	return t.tx.Commit()
}

func (t *oracleTx) Rollback(ctx context.Context) error {
	return t.tx.Rollback()
}
