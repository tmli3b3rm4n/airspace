package util

import "testing"

// Test cases for IsValidLatitude
func TestIsValidLatitude(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"Valid Latitude 0", 0.0, true},
		{"Valid Latitude Positive", 45.0, true},
		{"Valid Latitude Negative", -45.0, true},
		{"Valid Max Latitude", 90.0, true},
		{"Valid Min Latitude", -90.0, true},
		{"Invalid Latitude Over Max", 90.1, false},
		{"Invalid Latitude Under Min", -90.1, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsValidLatitude(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v for input %v, got %v", tc.expected, tc.input, result)
			}
		})
	}
}

// Test cases for IsValidLongitude
func TestIsValidLongitude(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"Valid Longitude 0", 0.0, true},
		{"Valid Longitude Positive", 100.0, true},
		{"Valid Longitude Negative", -100.0, true},
		{"Valid Max Longitude", 180.0, true},
		{"Valid Min Longitude", -180.0, true},
		{"Invalid Longitude Over Max", 180.1, false},
		{"Invalid Longitude Under Min", -180.1, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsValidLongitude(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v for input %v, got %v", tc.expected, tc.input, result)
			}
		})
	}
}
