package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestRestrictedAirspace(t *testing.T) {
	// Create a mock database
	DB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer DB.Close()

	// Wrap the mock DB in GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: DB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open GORM DB: %v", err)
	}

	repo := &flightRestrictions.FlightRestrictionsRepo{DB: gormDB}

	lat, lon := 40.7128, -74.0060
	mock.ExpectQuery(`SELECT count\(\*\) FROM "flight_restrictions" WHERE ST_Intersects\(geom, ST_SetSRID\(ST_MakePoint\(\$1, \$2\), 4326\)\)`).
		WithArgs(lon, lat).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1)) // Simulate 1 match

	result, err := repo.RestrictedAirspace(lat, lon)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Fatalf("expected true, got false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet mock expectations: %v", err)
	}
}
