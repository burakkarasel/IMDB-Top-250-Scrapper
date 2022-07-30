package controller

import (
	"strconv"

	"github.com/burakkarasel/IMDB-Top-250-Scrapper/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetMovieByID(c *fiber.Ctx) error {

	id := c.Params("id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if ID > 250 {
		return c.Status(404).SendString("404 Not Found")
	}

	movie, err := models.GetMovieById(ID)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")

	return c.JSON(&movie)
}

// GetAllMovies gets all the movies from the DB
func GetAllMovies(c *fiber.Ctx) error {

	movies, err := models.GetAllMovies()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")

	return c.JSON(&movies)
}
