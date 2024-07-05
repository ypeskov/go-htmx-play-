package routes

import (
	"Tpl/models"
	"Tpl/templates/pages"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TodoItemView struct {
	Id   int
	Item string
	Done bool
}

func (r *Routes) RegisterItemsRoutes(g *echo.Group) {
	g.POST("/add", r.AddItem)
	g.DELETE("/delete/:id", r.DeleteItem)
	g.PUT("/change-status/:id", r.ChangeItemStatus)
}

func (r *Routes) AddItem(c echo.Context) error {
	r.log.Info("Adding item to the database")

	item := models.TodoItem{}
	if err := c.Bind(&item); err != nil {
		r.log.Errorf("Failed to bind item: %v", err)
		return c.JSON(http.StatusBadRequest, "Failed to bind item")
	}

	if err := r.itemsService.AddItem(item); err != nil {
		r.log.Errorf("Failed to add item to the database: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to add item to the database")
	}

	items, err := r.itemsService.GetItems()
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}

	itemsListComp := pages.ItemsList(items)

	return Render(c, http.StatusOK, itemsListComp)
}

func (r *Routes) DeleteItem(c echo.Context) error {
	r.log.Info("Deleting item from the database")

	id := c.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		r.log.Errorf("Failed to parse id: %v", err)
		return c.JSON(http.StatusBadRequest, "Failed to parse id")
	}

	if err := r.itemsService.DeleteItem(idInt64); err != nil {
		r.log.Errorf("Failed to delete item from the database: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to delete item from the database")
	}

	return c.String(http.StatusOK, "")
}

func (r *Routes) ChangeItemStatus(c echo.Context) error {
	r.log.Info("Changing item status in the database")

	id := c.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		r.log.Errorf("Failed to parse id: %v", err)
		return c.JSON(http.StatusBadRequest, "Failed to parse id")
	}

	newStatusStr := c.QueryParam("newStatus")
	var newStatus bool
	if newStatusStr == "" || newStatusStr == "false" {
		newStatus = false
	} else {
		newStatus = true
	}

	if err := r.itemsService.ChangeItemStatus(idInt64, newStatus); err != nil {
		r.log.Errorf("Failed to change item status in the database: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to change item status in the database")
	}

	return c.String(http.StatusOK, "")
}

func (r *Routes) ShowIndexPage(c echo.Context) error {

	items, err := r.itemsService.GetItems()
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}

	itemsListComp := pages.ItemsList(items)
	return Render(c, http.StatusOK, pages.Home("fuck", itemsListComp))
}
