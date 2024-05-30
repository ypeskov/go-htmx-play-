package sqlite

import (
	"Tpl/internal/database"
	"Tpl/internal/logger"
	"Tpl/models"
)

type ItemRepository struct {
	log *logger.Logger
	db  *database.Database
}

type ItemRepositoryInterface interface {
	GetItems() ([]models.TodoItem, error)
}

func New(log *logger.Logger, db *database.Database) ItemRepositoryInterface {
	return &ItemRepository{
		log: log,
		db:  db,
	}
}

func (r *ItemRepository) GetItems() ([]models.TodoItem, error) {
	query := "SELECT id, item, IFNULL(NULLIF(done, ''), 'false') AS done FROM todos"
	var items []models.TodoItem
	err := r.db.Select(&items, query)
	if err != nil {
		r.log.Errorf("REPO: Failed to execute query: %v", err)
		return nil, err
	}

	return items, nil
}
