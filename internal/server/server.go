package server

import (
	"Tpl/internal/config"
	"Tpl/internal/database"
	"Tpl/internal/logger"
	"Tpl/internal/routes"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e    *echo.Echo
	port string
	log  *logger.Logger
}

func New(cfg *config.Config,
	log *logger.Logger,
	db *database.Database) *Server {
	e := echo.New()

	e.HideBanner = true

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	handlers := routes.New(log, db)

	itemsGroup := e.Group("/items")
	handlers.RegisterItemsRoutes(itemsGroup)

	return &Server{
		e:    e,
		port: cfg.Port,
		log:  log,
	}

}

func (s *Server) Start() error {
	s.log.Infof("Starting the server on port %s", s.port)

	return s.e.Start(s.port)
}
