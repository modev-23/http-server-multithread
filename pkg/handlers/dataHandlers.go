package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/modev-23/http-multithread-caching/config"
	"github.com/modev-23/http-multithread-caching/internal/models"
)

// Handler
func GetMovie(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var movie models.Rating

	// Query the database using raw SQL
	row := db.QueryRow("SELECT * FROM movies WHERE rating_id = $1", id)
	if err := row.Scan(&movie.RatingId, &movie.UserId, &movie.MovieId, &movie.Rating); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Movie not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error retrieving movie",
			"error":   err.Error(),
		})
	}

	response := map[string]interface{}{
		"data": movie,
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllMovies(c echo.Context) error {
	db := config.DB()

	// Query the database using raw SQL
	rows, err := db.QueryContext(context.Background(), "SELECT * FROM movies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to hold the movies
	var movies []models.Rating

	// Iterate through the rows
	for rows.Next() {
		var movie models.Rating
		err := rows.Scan(&movie.RatingId, &movie.UserId, &movie.MovieId, &movie.Rating)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	// Check for any errors during row iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the movies slice as JSON response
	return c.JSON(http.StatusOK, movies)
}
