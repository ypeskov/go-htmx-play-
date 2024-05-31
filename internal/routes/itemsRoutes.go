package routes

import (
	"Tpl/models"
	"bytes"
	"github.com/CloudyKit/jet"
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
	g.GET("/all", r.GetItems)
	g.POST("/add", r.AddItem)
	g.DELETE("/delete/:id", r.DeleteItem)
	g.PUT("/change-status/:id", r.ChangeItemStatus)
}

func (r *Routes) GetItems(c echo.Context) error {
	//r.log.Info("Getting items from the database")

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

	if err := r.itemsService.AddItem(item); err != nil {
		r.log.Errorf("Failed to add item to the database: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to add item to the database")
	}

	items, err := r.itemsService.GetItems()
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}

	t, err := r.View.GetTemplate("components/items-list.jet")
	if err != nil {
		r.log.Errorf("Failed to parse template: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to parse template")
	}

	vars := make(jet.VarMap)
	convertedItems := convertItemsToView(items)
	vars.Set("items", convertedItems)

	var buf bytes.Buffer
	if err := t.Execute(&buf, vars, convertedItems); err != nil {
		r.log.Errorf("Failed to execute template: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to execute template")
	}

	return c.HTMLBlob(http.StatusOK, buf.Bytes())
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
	t, err := r.View.GetTemplate("index.jet")
	if err != nil {
		r.log.Errorf("Failed to get template: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to get template")
	}
	r.View.SetDevelopmentMode(true)

	items, err := r.itemsService.GetItems()
	if err != nil {
		r.log.Errorf("Failed to get items from the database: %v", err)
		return c.JSON(500, "Failed to get items from the database")
	}

	data := make(jet.VarMap)
	data.Set("items", convertItemsToView(items))

	var buf bytes.Buffer
	if err := t.Execute(&buf, data, nil); err != nil {
		r.log.Errorf("Failed to render index page: %v", err)
		return c.JSON(500, "Failed to render index page")
	}

	return c.HTMLBlob(http.StatusOK, buf.Bytes())
}

func convertItemsToView(items []models.TodoItem) []TodoItemView {
	var itemsView []TodoItemView
	for _, item := range items {
		itemsView = append(itemsView, TodoItemView{
			Id:   item.Id,
			Item: item.Item,
			Done: *item.Done,
		})
	}
	return itemsView
}
