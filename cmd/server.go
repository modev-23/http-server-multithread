package main

import (
	"github.com/labstack/echo/v4"
	"github.com/modev-23/http-multithread-caching/config"
	"github.com/modev-23/http-multithread-caching/pkg/handlers"
)

func main() {
	e := echo.New()

	movieRoute := e.Group("/movie")
	movieRoute.GET("/:id", handlers.GetMovie)
	movieRoute.GET("/all", handlers.GetAllMovies)

	config.DatabaseInit()

	// Ping the database to verify the connection
	db := config.DB()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
