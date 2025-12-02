package db

import (
	"database/sql"
)

// ============= STRUCT =============
type OracleClient struct {
	DB *sql.DB
}

type oracleTx struct {
	tx *sql.Tx
}
