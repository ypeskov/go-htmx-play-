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
