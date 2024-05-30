package routes

import (
	"Tpl/models"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"strconv"
)

func (r *Routes) RegisterItemsRoutes(g *echo.Group) {
	g.GET("/all", r.GetItems)
	g.POST("/add", r.AddItem)
	g.DELETE("/delete/:id", r.DeleteItem)
}

func (r *Routes) GetItems(c echo.Context) error {
	r.log.Info("Getting items from the database")

	items, err := r.itemsService.GetItems()
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}

	return c.JSON(200, items)
}

func (r *Routes) AddItem(c echo.Context) error {
	r.log.Info("Adding item to the database")

	item := models.TodoItem{}
	if err := c.Bind(&item); err != nil {
		r.log.Errorf("Failed to bind item: %v", err)
		return c.JSON(http.StatusBadRequest, "Failed to bind item")
	}
	r.log.Infof("Item: %+v", item)

	if err := r.itemsService.AddItem(item); err != nil {
		r.log.Errorf("Failed to add item to the database: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to add item to the database")
	}

	items, err := r.itemsService.GetItems()
	r.log.Infof("Items: %+v", items)
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}

	t, err := template.ParseFiles("templates/components/items-list.html")
	if err != nil {
		r.log.Errorf("Failed to parse template: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to parse template")
	}

	if err := t.ExecuteTemplate(c.Response(), "items", items); err != nil {
		r.log.Errorf("Failed to execute template: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to execute template")
	}

	return nil
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

func (r *Routes) ShowIndexPage(c echo.Context) error {
	t := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/index.html",
		"templates/components/items-list.html",
	))

	items, err := r.itemsService.GetItems()
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}
	r.log.Info("Items: %+v", items)

	if err := t.ExecuteTemplate(c.Response(), "index", items); err != nil {
		r.log.Errorf("Failed to render index page: %v", err)
		return c.JSON(500, "Failed to render index page")
	}

	return nil
}
