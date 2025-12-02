package routes

import (
	"belajar/app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/oracle/lol", handler.GetSavedNumbersWow)

}
