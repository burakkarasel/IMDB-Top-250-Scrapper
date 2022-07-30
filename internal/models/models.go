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

// GetMovieById gets the movie from DB according to it's id
func GetMovieById(id int) (Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select title, poster, genre, year, rating
		from imdb_movies
		where id = $1
	`
	var mv Movie

	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&mv.Title,
		&mv.Poster,
		&mv.Genre,
		&mv.Year,
		&mv.Rating,
	)

	if err != nil {
		return mv, err
	}

	return mv, nil
}

// GetAllMovies returns all the movies from DB as a slice of Movie
func GetAllMovies() ([]Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select title, poster, genre, year, rating
		from imdb_movies
	`
	var movies []Movie

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return movies, err
	}

	for rows.Next() {
		var mv Movie
		err := rows.Scan(
			&mv.Title,
			&mv.Poster,
			&mv.Genre,
			&mv.Year,
			&mv.Rating,
		)

		if err != nil {
			return movies, err
		}

		movies = append(movies, mv)
	}

	if err = rows.Err(); err != nil {
		return movies, err
	}

	return movies, nil
}
