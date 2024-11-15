package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/tmli3b3rm4n/airspace/internal/handler"
	"github.com/tmli3b3rm4n/airspace/internal/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/airspace/internal/database"
	"github.com/tmli3b3rm4n/airspace/pkg/formulas"
)

var restrictedAreas []RestrictedArea

// Check if the given latitude and longitude are within a restricted zone
func isRestricted(lat, lon float64) bool {
	for _, area := range restrictedAreas {
		distance := formulas.Haversine(lat, lon, area.Latitude, area.Longitude)
		if distance <= area.Radius {
			return true // Inside a restricted area
		}
	}
	return false // Not in a restricted area
}

// Handler for checking if drone flight is allowed
func checkFlightPermission(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")

	// Convert to float64
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	// Check if the location is restricted
	if isRestricted(latitude, longitude) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"message": "Flight not allowed"})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Flight allowed"})
	}
}

// Load restricted areas from CSV file
func loadRestrictedAreas(filename string) ([]RestrictedArea, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var areas []RestrictedArea
	for _, record := range records[1:] { // Skip header
		latitude, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		longitude, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		radius, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}

		areas = append(areas, RestrictedArea{
			Latitude:  latitude,
			Longitude: longitude,
			Radius:    radius,
		})
	}

	return areas, nil
}

func processGeoJSON(db *gorm.DB, jsonData []byte) error {
	var featureCollection database.FeatureCollection

	err := json.Unmarshal(jsonData, &featureCollection)
	if err != nil {
		return err
	}

	for _, feature := range featureCollection.Features {
		area := database.RestrictedArea{
			ObjectID:      feature.Properties.ObjectID,
			Proponent:     feature.Properties.Proponent,
			Branch:        feature.Properties.Branch,
			Base:          feature.Properties.Base,
			Facility:      feature.Properties.Facility,
			Airspace:      feature.Properties.Airspace,
			Reason:        feature.Properties.Reason,
			State:         feature.Properties.State,
			FAA_ID:        feature.Properties.FAA_ID,
			POC:           feature.Properties.POC,
			Floor:         feature.Properties.Floor,
			Ceiling:       feature.Properties.Ceiling,
			County:        feature.Properties.County,
			Shape__Area:   feature.Properties.Shape__Area,
			Shape__Length: feature.Properties.Shape__Length,
		}

		// Handle the Geometry (polygon coordinates)
		// Here's a simple example of storing the first and last points:
		polygon := feature.Geometry.Coordinates[0]
		area.SimplifiedGeometry = fmt.Sprintf("POLYGON((%f %f, %f %f))", polygon[0][0], polygon[0][1], polygon[len(polygon)-1][0], polygon[len(polygon)-1][1])

		// Insert the area into the database
		if err := db.Create(&area).Error; err != nil {
			return err
		}
	}

	return nil
}

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Connected to airspace database")

	err = db.AutoMigrate()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Migrated airspace database")

	var restrictedArea database.RestrictedArea
	if err := db.First(&restrictedArea); err == nil {
		err := loadRecords("National_Security_UAS_Flight_Restrictions.csv", db)
		if err != nil {
			log.Printf("Error loading restricted areas: %v", err)
			return
		}
	}

	e := echo.New()

	restrictedAreaRepo := repository.NewRestrictedAreaRepository(db)
	restrictedAreaHandler := handler.NewRestrictedAreaHandler(restrictedAreaRepo)
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://localhost:3000", "https://www.gitasy.com", "https://gitasy.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/is-restricted/:lat/:long", restrictedAreaHandler)

	// Setup API routes
	http.HandleFunc("/check-flight", checkFlightPermission)

	// Start server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
