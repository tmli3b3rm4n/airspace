package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Define database connection details
	host := os.Getenv("POSTGRES_HOST")
	port := 5432
	dbName := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	// Define the GeoJSON file to load
	geojsonFile := "/data/National_Security_UAS_Flight_Restrictions.geojson"
	log.Printf("User: %v", user)
	// Construct ogr2ogr command
	cmd := exec.Command(
		"ogr2ogr",
		"-f", "PostgreSQL",
		fmt.Sprintf("PG:host=%s port=%d dbname=%s user=%s password=%s", host, port, dbName, user, password),
		geojsonFile,
		"-nln", "flight_restrictions",
		"-lco", "GEOMETRY_NAME=geom",
		"-overwrite", // Overwrite existing data
	)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return
	}

	fmt.Println("Data successfully loaded.")
}
