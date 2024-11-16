package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tmli3b3rm4n/airspace/internal/database"
	"github.com/tmli3b3rm4n/airspace/internal/handler"
	"github.com/tmli3b3rm4n/airspace/internal/repository"
	"log"
)

func main() {
	// Establish database connection
	connect, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Initialize Echo framework
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize repository and handler
	flightRestrictionRepo := repository.NewFlightRestrictionsRepo(connect)
	flightRestrictionHandler := handler.NewFlightRestrictionsHandler(flightRestrictionRepo)

	// Register the GET route
	e.GET("/restricted-airspace/:lat/:lon", flightRestrictionHandler.RestrictedAirspace)

	// Start Echo server on port 8080
	fmt.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
