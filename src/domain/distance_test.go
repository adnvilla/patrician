package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	// Test creating a distance instance
	distance := domain.Distance{
		FromCity: "Estocolmo",
		ToCity:   "Visby",
		Value:    15.5,
	}

	// Test field access
	assert.Equal(t, "Estocolmo", distance.FromCity)
	assert.Equal(t, "Visby", distance.ToCity)
	assert.Equal(t, float32(15.5), distance.Value)

	// Test that it embeds Entity
	assert.Equal(t, uint(0), distance.ID) // Initial ID should be zero
	assert.NotPanics(t, func() {
		_ = distance.CreatedAt
		_ = distance.UpdatedAt
		_ = distance.DeletedAt
	})
}

func TestDistanceModification(t *testing.T) {
	// Test modifying distance fields
	distance := domain.Distance{
		FromCity: "City1",
		ToCity:   "City2",
		Value:    10.0,
	}

	// Modify fields
	distance.FromCity = "NewCity1"
	distance.ToCity = "NewCity2"
	distance.Value = 25.5

	// Verify modifications
	assert.Equal(t, "NewCity1", distance.FromCity)
	assert.Equal(t, "NewCity2", distance.ToCity)
	assert.Equal(t, float32(25.5), distance.Value)
}

func TestDistanceZeroValue(t *testing.T) {
	// Test zero value of Distance
	var distance domain.Distance

	assert.Equal(t, "", distance.FromCity)
	assert.Equal(t, "", distance.ToCity)
	assert.Equal(t, float32(0), distance.Value)
	assert.Equal(t, uint(0), distance.ID)
}