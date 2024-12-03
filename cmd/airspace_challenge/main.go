package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/tmli3b3rm4n/airspace/docs"
	"github.com/tmli3b3rm4n/airspace/internal/database"
	"github.com/tmli3b3rm4n/airspace/internal/handler"
	"github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions"
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	fmt.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
