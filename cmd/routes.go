package main

import (
	"github.com/bruce-mig/trivia-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.GetFact)

	app.Delete("/fact/:id", handlers.DeleteFact)

	app.Put("/fact/:id", handlers.UpdateFact)
}
