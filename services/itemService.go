package services

import (
	"Tpl/internal/database"
	"Tpl/internal/logger"
	"Tpl/models"
	repository "Tpl/reporitories/sqlite"
)

type ItemService struct {
	log      *logger.Logger
	db       *database.Database
	itemRepo repository.ItemRepositoryInterface
}

type ItemServiceInterface interface {
	GetItems() ([]models.TodoItem, error)
	AddItem(item models.TodoItem) error
	DeleteItem(id int64) error
}

func NewItemsService(logger *logger.Logger,
	db *database.Database,
	itemRepository repository.ItemRepositoryInterface) ItemServiceInterface {

	return &ItemService{
		log:      logger,
		db:       db,
		itemRepo: itemRepository,
	}
}

func (s *ItemService) GetItems() ([]models.TodoItem, error) {
	return s.itemRepo.GetItems()
}

func (s *ItemService) AddItem(item models.TodoItem) error {
	return s.itemRepo.AddItem(item)
}

func (s *ItemService) DeleteItem(id int64) error {
	return s.itemRepo.DeleteItem(id)
}
