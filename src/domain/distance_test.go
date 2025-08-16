package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

// Distance structure tests
func TestDistanceStructure(t *testing.T) {
	distance := domain.Distance{
		FromCity: "Estocolmo",
		ToCity:   "Visby",
		Value:    1.5,
	}

	assert.Equal(t, "Estocolmo", distance.FromCity)
	assert.Equal(t, "Visby", distance.ToCity)
	assert.Equal(t, float32(1.5), distance.Value)

	// Verify Entity inheritance
	assert.Equal(t, uint(0), distance.ID)
	assert.True(t, distance.CreatedAt.IsZero())
}

func TestDistanceEntityInheritance(t *testing.T) {
	distance := domain.Distance{
		FromCity: "TestCityA",
		ToCity:   "TestCityB",
		Value:    2.5,
	}

	// Test Entity fields
	assert.Equal(t, uint(0), distance.Entity.ID)
	assert.True(t, distance.Entity.CreatedAt.IsZero())
	assert.True(t, distance.Entity.UpdatedAt.IsZero())
	assert.False(t, distance.Entity.DeletedAt.Valid)
}

func TestDistanceZeroValue(t *testing.T) {
	distance := domain.Distance{}

	assert.Equal(t, "", distance.FromCity)
	assert.Equal(t, "", distance.ToCity)
	assert.Equal(t, float32(0), distance.Value)
}

func TestDistanceNegativeValue(t *testing.T) {
	distance := domain.Distance{
		FromCity: "CityA",
		ToCity:   "CityB",
		Value:    -1.0,
	}

	assert.Equal(t, float32(-1.0), distance.Value)
}

func TestDistanceFloatPrecision(t *testing.T) {
	distance := domain.Distance{
		FromCity: "CityA",
		ToCity:   "CityB",
		Value:    1.23456789,
	}

	// float32 has limited precision
	assert.InDelta(t, 1.23456789, distance.Value, 0.000001)
}

// Distances map tests
func TestDistancesMapStructure(t *testing.T) {
	distances := domain.Distances

	assert.NotNil(t, distances)
	assert.Greater(t, len(distances), 0, "Distances map should not be empty")

	// Test that all cities in the Cities map have distance entries
	for cityName := range domain.Cities {
		cityDistances, exists := distances[cityName]
		assert.True(t, exists, "City %s should have distance entries", cityName)
		assert.NotNil(t, cityDistances, "Distance map for %s should not be nil", cityName)
	}
}

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

func TestDistancesCompleteness(t *testing.T) {
	distances := domain.Distances

	// Count expected cities
	expectedCities := len(domain.Cities)

	// Each city should have distances to all other cities (excluding itself)
	for fromCity, fromDistances := range distances {
		// Should have distances to all other cities
		expectedDistanceCount := expectedCities - 1 // Excluding itself
		assert.Equal(t, expectedDistanceCount, len(fromDistances),
			"City %s should have distances to %d other cities", fromCity, expectedDistanceCount)

		// Should not have distance to itself
		_, hasSelfDistance := fromDistances[fromCity]
		assert.False(t, hasSelfDistance, "City %s should not have distance to itself", fromCity)
	}
}

func TestDistancesPositiveValues(t *testing.T) {
	distances := domain.Distances

	// All distances should be positive
	for fromCity, fromDistances := range distances {
		for toCity, distance := range fromDistances {
			assert.Greater(t, distance, float32(0),
				"Distance from %s to %s should be positive", fromCity, toCity)
		}
	}
}

func TestDistancesSpecificValues(t *testing.T) {
	distances := domain.Distances

	// Test some specific known distances
	if estocolmoDistances, exists := distances["Estocolmo"]; exists {
		if visbyDistance, exists := estocolmoDistances["Visby"]; exists {
			assert.Greater(t, visbyDistance, float32(0), "Distance from Estocolmo to Visby should be positive")
		}
	}
}

func TestDistancesCityConsistency(t *testing.T) {
	distances := domain.Distances
	cities := domain.Cities

	// All cities in distances should exist in Cities map
	for cityName := range distances {
		_, exists := cities[cityName]
		assert.True(t, exists, "City %s from distances should exist in Cities map", cityName)
	}

	// All cities in Cities map should have distance entries
	for cityName := range cities {
		_, exists := distances[cityName]
		assert.True(t, exists, "City %s from Cities should have distance entries", cityName)
	}
}

func TestDistancesDataIntegrity(t *testing.T) {
	distances := domain.Distances

	for fromCity, fromDistances := range distances {
		for toCity, distance := range fromDistances {
			// Distance should be a valid number
			assert.False(t, distance != distance, "Distance from %s to %s should not be NaN", fromCity, toCity) // NaN check
			assert.False(t, distance == 0, "Distance from %s to %s should not be zero", fromCity, toCity)
			assert.Greater(t, distance, float32(0), "Distance from %s to %s should be positive", fromCity, toCity)
		}
	}
}
