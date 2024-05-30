package routes

import (
	"Tpl/internal/database"
	"Tpl/internal/logger"
	"Tpl/reporitories/sqlite"
	"Tpl/services"
)

type Routes struct {
	log          *logger.Logger
	itemsService services.ItemServiceInterface
}

func New(log *logger.Logger, db *database.Database) *Routes {
	itemRepo := sqlite.New(log, db)
	itemsService := services.NewItemsService(log, db, itemRepo)

	return &Routes{
		log:          log,
		itemsService: itemsService,
	}
}
