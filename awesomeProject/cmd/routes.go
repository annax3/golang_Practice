package main

import (
	"AwesomeProjects/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/listFact", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
