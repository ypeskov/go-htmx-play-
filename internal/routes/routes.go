package routes

import (
	"Tpl/internal/database"
	"Tpl/internal/logger"
	"Tpl/reporitories/sqlite"
	"Tpl/services"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	log          *logger.Logger
	itemsService services.ItemServiceInterface
}

func New(log *logger.Logger, db *database.Database) *Routes {
	itemRepo := sqlite.New(log, db)
	itemsService := services.NewItemsService(log, db, itemRepo)

	return &Routes{
		log:          log,
		itemsService: itemsService,
	}
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
