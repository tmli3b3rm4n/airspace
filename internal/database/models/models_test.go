package models

import (
	"database/sql/driver"
	"testing"
)

func TestPostgisGeometry_Scan(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		expectErr bool
		expected  string
	}{
		{
			name:      "valid string input",
			input:     "POINT(30 10)",
			expectErr: false,
			expected:  "POINT(30 10)",
		},
		{
			name:      "invalid input type",
			input:     12345,
			expectErr: true,
			expected:  "",
		},
		{
			name:      "nil input",
			input:     nil,
			expectErr: true,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var geom PostgisGeometry
			err := geom.Scan(tt.input)

			if tt.expectErr && err == nil {
				t.Errorf("expected an error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("did not expect an error but got: %v", err)
			}
			if geom.Geometry != tt.expected {
				t.Errorf("expected geometry %s, got %s", tt.expected, geom.Geometry)
			}
		})
	}
}

func TestPostgisGeometry_Value(t *testing.T) {
	tests := []struct {
		name      string
		input     PostgisGeometry
		expectErr bool
		expected  driver.Value
	}{
		{
			name:      "valid geometry",
			input:     PostgisGeometry{Geometry: "POLYGON((30 10, 40 40, 20 40, 10 20, 30 10))"},
			expectErr: false,
			expected:  "POLYGON((30 10, 40 40, 20 40, 10 20, 30 10))",
		},
		{
			name:      "empty geometry",
			input:     PostgisGeometry{Geometry: ""},
			expectErr: false,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := tt.input.Value()

			if tt.expectErr && err == nil {
				t.Errorf("expected an error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("did not expect an error but got: %v", err)
			}
			if value != tt.expected {
				t.Errorf("expected value %v, got %v", tt.expected, value)
			}
		})
	}
}

func TestFlightRestriction_Struct(t *testing.T) {
	fr := FlightRestriction{
		ID:          1,
		Proponent:   "Test Proponent",
		Branch:      "Air Force",
		Base:        "Test Base",
		Facility:    "Test Facility",
		Airspace:    "Test Airspace",
		Reason:      "Test Reason",
		State:       "TX",
		FaaID:       "TX123",
		POC:         "Test POC",
		Floor:       1000,
		Ceiling:     2000,
		County:      "Test County",
		ShapeArea:   12345.67,
		ShapeLength: 89.01,
		Geom:        PostgisGeometry{Geometry: "POINT(30 10)"},
	}

	if fr.ID != 1 {
		t.Errorf("expected ID 1, got %d", fr.ID)
	}
	if fr.Proponent != "Test Proponent" {
		t.Errorf("expected Proponent 'Test Proponent', got %s", fr.Proponent)
	}
	if fr.Geom.Geometry != "POINT(30 10)" {
		t.Errorf("expected Geom 'POINT(30 10)', got %s", fr.Geom.Geometry)
	}
}
