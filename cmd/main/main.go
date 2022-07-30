package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/burakkarasel/IMDB-Top-250-Scrapper/internal/models"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func main() {

	tempMv, _ := models.GetMovieById(1)

	if tempMv.Title == "" {
		topMovies := crawl()
		for _, m := range topMovies {
			err := models.InsertMovie(m)

			if err != nil {
				log.Fatal("cannot insert value to DB")
			}
		}
	}

	app = fiber.New()
	Router(app)

	app.Listen(":3000")
}

// crawl goes to IMDB top250 page and scraps movie data for each movie's page
func crawl() []models.Movie {
	var mvSlc []models.Movie
	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)

	infoCollector := c.Clone()

	c.OnHTML(".titleColumn", func(e *colly.HTMLElement) {
		movieUrl := e.ChildAttr("a", "href")
		movieUrl = e.Request.AbsoluteURL(movieUrl)
		infoCollector.Visit(movieUrl)
	})

	infoCollector.OnHTML("section.sc-c7f03a63-0", func(e *colly.HTMLElement) {

		tmpMovie := models.Movie{}
		tmpMovie.Title = e.ChildText("div.khmuXj > h1")                                                      // done
		tmpMovie.Genre = e.ChildText("a.sc-16ede01-3 > span.ipc-chip__text")                                 // done
		tmpMovie.Poster = e.ChildAttr("div.ipc-media--poster-27x40 > img.ipc-image", "src")                  // done
		rating := e.ChildText("div[data-testid=\"hero-rating-bar__aggregate-rating__score\"] > span.jGRxWM") // not done
		rating = string(rating[:3])

		tmpMovie.Rating, _ = strconv.ParseFloat(rating, 64)

		year := e.ChildText("li.ipc-inline-list__item > span.sc-8c396aa2-2")
		year = string(year[:4])

		tmpMovie.Year, _ = strconv.Atoi(year)

		mvSlc = append(mvSlc, tmpMovie)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting:", r.URL.String())
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting profile URL:", r.URL.String())
	})

	startURL := "https://www.imdb.com/chart/top/"
	c.Visit(startURL)

	return mvSlc
}
