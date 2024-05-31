package routes

import (
	"Tpl/internal/database"
	"Tpl/internal/logger"
	"Tpl/reporitories/sqlite"
	"Tpl/services"
	"github.com/CloudyKit/jet"
	"os"
	"path/filepath"
)

type Routes struct {
	log          *logger.Logger
	itemsService services.ItemServiceInterface
	View         *jet.Set
}

func New(log *logger.Logger, db *database.Database) *Routes {
	itemRepo := sqlite.New(log, db)
	itemsService := services.NewItemsService(log, db, itemRepo)

	var root, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}

	return &Routes{
		log:          log,
		itemsService: itemsService,
		View:         jet.NewHTMLSet(filepath.Join(root, "templates")),
	}
}
