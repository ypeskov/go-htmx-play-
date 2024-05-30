package database

import (
	"Tpl/internal/config"
	log "Tpl/internal/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database = sqlx.DB

func New(cfg *config.Config, logger *log.Logger) (*Database, error) {
	db, err := sqlx.Open("sqlite3", cfg.DatabaseUrl)
	if err != nil {
		logger.Errorf("Failed to open the database: %v", err)

		return nil, err
	}

	return db, nil
}
