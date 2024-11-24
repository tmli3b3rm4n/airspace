package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	host := os.Getenv("POSTGRES_HOST")
	port := 5432
	dbName := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	geojsonFile := "/data/National_Security_UAS_Flight_Restrictions.geojson"
	log.Printf("User: %v", user)
	cmd := exec.Command(
		"ogr2ogr",
		"-f", "PostgreSQL",
		fmt.Sprintf("PG:host=%s port=%d dbname=%s user=%s password=%s", host, port, dbName, user, password),
		geojsonFile,
		"-nln", "flight_restrictions",
		"-lco", "GEOMETRY_NAME=geom",
		"-overwrite",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return
	}

	fmt.Println("Data successfully loaded.")
}
