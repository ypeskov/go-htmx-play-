package routes

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

func (r *Routes) RegisterItemsRoutes(g *echo.Group) {
	g.GET("/all", r.GetItems)
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

func (r *Routes) ShowIndexPage(c echo.Context) error {
	t := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/index.html",
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
