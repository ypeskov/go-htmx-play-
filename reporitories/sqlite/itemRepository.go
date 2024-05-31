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
	AddItem(item models.TodoItem) error
	DeleteItem(id int64) error
	ChangeItemStatus(id int64, newStatus bool) error
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
		r.log.Errorf("REPO: Failed to execute query: %v\n", err)
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) AddItem(item models.TodoItem) error {
	query := "INSERT INTO todos (item, done) VALUES (?, ?)"
	_, err := r.db.Exec(query, item.Item, item.Done)
	if err != nil {
		r.log.Errorf("REPO: Failed to execute query: %v\n", err)
		return err
	}

	return nil
}

func (r *ItemRepository) DeleteItem(id int64) error {
	query := "DELETE FROM todos WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Errorf("REPO: Failed to execute query: %+v\n", err)
		return err
	}

	return nil
}

func (r *ItemRepository) ChangeItemStatus(id int64, newStatus bool) error {
	r.log.Infof("REPO: Changing status of item with id: %d to: %t\n", id, newStatus)
	query := "UPDATE todos SET done = ? WHERE id = ?"
	result, err := r.db.Exec(query, newStatus, id)
	if err != nil {
		r.log.Errorf("REPO: Failed to execute query: %+v\n", err)
		return err
	}
	rowsQty, _ := result.RowsAffected()
	r.log.Infof("REPO: Rows affected: %d\n", rowsQty)

	return nil
}
