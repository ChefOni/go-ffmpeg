package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
	"fmt"
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


func timeConversion(s string) string {
    // Write your code here

    parsedTime, err := time.Parse("03:04 PM", s)
    if err != nil {
        fmt.Println("Error parsing time:", err)
        return ""
    }
    fmt.Println(parsedTime.String())
    return parsedTime.GoString()
}
