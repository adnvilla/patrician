package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestDistanceRepositoryAdditional(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewDistanceRepository(db)

	// Create multiple test distances to verify FindAll functionality
	distances := []domain.Distance{
		{FromCity: "TestCity1", ToCity: "TestCity2", Value: 15.5},
		{FromCity: "TestCity2", ToCity: "TestCity3", Value: 25.0},
		{FromCity: "TestCity1", ToCity: "TestCity3", Value: 35.5},
	}

	// Test Create for multiple distances
	for i := range distances {
		err := repo.Create(&distances[i])
		assert.NoError(t, err)
		assert.NotEqual(t, uint(0), distances[i].ID)
	}

	// Test FindAll returns all created distances
	found, err := repo.FindAll()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(found), len(distances))

	// Test Update functionality
	distances[0].Value = 20.5
	err = repo.Update(&distances[0])
	assert.NoError(t, err)

	// Verify the update worked by checking FindAll again
	updated, err := repo.FindAll()
	assert.NoError(t, err)
	
	// Find our updated distance and verify the value
	var foundUpdated bool
	for _, d := range updated {
		if d.ID == distances[0].ID {
			assert.Equal(t, float32(20.5), d.Value)
			foundUpdated = true
			break
		}
	}
	assert.True(t, foundUpdated, "Updated distance should be found")

	// Test Delete functionality
	for _, distance := range distances {
		err = repo.Delete(&distance)
		assert.NoError(t, err)
	}
}
