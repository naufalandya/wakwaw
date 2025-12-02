package db

import (
	"belajar/app/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var Oracle *sql.DB // GLOBAL ORACLE CLIENT

func InitOracle(cfg config.OracleConfig) error {
	dsn := fmt.Sprintf(
		`user="%s" password="%s" connectString="%s:%s/%s" timezone=Asia/Jakarta`,
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Service,
	)

	if cfg.Role == "SYSDBA" {
		dsn += " sysdba=1"
	}

	db, err := sql.Open("godror", dsn)
	if err != nil {
		return fmt.Errorf("failed to open Oracle connection: %w", err)
	}

	// pool
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("oracle ping failed: %w", err)
	}

	log.Println("[DB] Oracle connected successfully")
	Oracle = db
	return nil
}
