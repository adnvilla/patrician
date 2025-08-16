package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestDistancesConsistency(t *testing.T) {
	// Test that all cities in the distances map have entries for all other cities
	for fromCity, distancesMap := range domain.Distances {
		for toCity := range domain.Distances {
			if fromCity != toCity {
				// Check that a distance exists from fromCity to toCity
				distance, exists := distancesMap[toCity]
				assert.True(t, exists, "Distance from %s to %s should exist", fromCity, toCity)
				assert.Greater(t, distance, float32(0), "Distance from %s to %s should be positive", fromCity, toCity)
			}
		}
	}
}

func TestDistancesSymmetry(t *testing.T) {
	// Test that distances are symmetric (distance from A to B equals distance from B to A)
	for fromCity, distancesMap := range domain.Distances {
		for toCity, distance := range distancesMap {
			if fromCity != toCity {
				// Get the reverse distance
				reverseDistance, exists := domain.Distances[toCity][fromCity]
				assert.True(t, exists, "Reverse distance from %s to %s should exist", toCity, fromCity)
				assert.Equal(t, distance, reverseDistance,
					"Distance from %s to %s should equal distance from %s to %s",
					fromCity, toCity, toCity, fromCity)
			}
		}
	}
}

func TestDistancesWithCities(t *testing.T) {
	// Test that all cities in the distances map exist in the Cities map
	for cityName := range domain.Distances {
		_, exists := domain.Cities[cityName]
		assert.True(t, exists, "City %s in distances map should exist in Cities map", cityName)
	}

	// Test that all cities in the Cities map exist in the distances map
	for cityName := range domain.Cities {
		_, exists := domain.Distances[cityName]
		assert.True(t, exists, "City %s in Cities map should exist in distances map", cityName)
	}
}
