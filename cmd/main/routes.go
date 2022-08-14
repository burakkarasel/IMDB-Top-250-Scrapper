package main

import (
	"github.com/burakkarasel/IMDB-Top-250-Scrapper/internal/controller"
	"github.com/gofiber/fiber/v2"
)

// Router holds the routes for the app
func Router(app *fiber.App) {
	app.Get("/api/all", controller.GetAllMovies)
	app.Get("/api/:id", controller.GetMovieByID)
}
