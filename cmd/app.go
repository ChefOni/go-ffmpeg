package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	router *fiber.App
	logger *log.Logger
	db     *gorm.DB
}

func New(logger *log.Logger, db *gorm.DB) *App {
	app := fiber.New() 

	return &App{
		logger: logger,
		router: app,
		db:     db,
	}
}

func (a *App) Start(address string) error {
	a.logger.Println("Starting server on", address)
	return a.router.Listen(address)
}

