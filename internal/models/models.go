package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/burakkarasel/IMDB-Top-250-Scrapper/internal/config"
)

// Movie holds the movie data
type Movie struct {
	Title  string  `json:"title"`
	Poster string  `json:"poster"`
	Genre  string  `json:"genre"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

var db *sql.DB

// init creates DB connection
func init() {
	config.Connect()
	db = config.GetDB()
}

// InsertMovie inserts a movie into DB
func InsertMovie(m Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into imdb_movies (title, poster, genre, year,
		rating, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.ExecContext(ctx, query, m.Title, m.Poster, m.Genre, m.Year, m.Rating, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
