package routes

import "github.com/labstack/echo/v4"

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
