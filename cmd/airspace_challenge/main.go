package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tmli3b3rm4n/airspace/internal/database"
	"github.com/tmli3b3rm4n/airspace/internal/handler"
	"github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions"
	"log"
)

func main() {
	connect, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	flightRestrictionRepo := flightRestrictions.NewFlightRestrictionsRepo(connect)
	flightRestrictionHandler := handler.NewFlightRestrictionsHandler(flightRestrictionRepo)

	e.GET("/restricted-airspace/:lat/:lon", flightRestrictionHandler.RestrictedAirspace)

	fmt.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
