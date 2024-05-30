package main

import (
	"Tpl/internal/config"
	"Tpl/internal/database"
	log "Tpl/internal/logger"
	"Tpl/internal/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := log.New(cfg)
	logger.Info("Starting the application...")

	db, err := database.New(cfg, logger)
	if err != nil {
		logger.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func(db *database.Database) {
		err := db.Close()
		if err != nil {
			logger.Fatalf("Failed to close the database connection: %v", err)
		}
	}(db)

	srv := server.New(cfg, logger, db)
	if err := srv.Start(); err != nil {
		logger.Errorf("Failed to start the server: %v", err)
	}

}
