package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestCityRepositoryAdditional(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCityRepository(db)

	// Create test city with market hall
	commodities := domain.GetCommodities()
	marketHall := domain.MarketHall{Commodities: commodities}
	city := domain.City{
		Name:       "TestCity",
		MarketHall: marketHall,
	}

	// Test Create
	err := repo.Create(&city)
	assert.NoError(t, err)
	assert.NotEqual(t, uint(0), city.ID)

	// Test FindAll
	cities, err := repo.FindAll()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(cities), 1)

	// Verify the city exists in the results
	var found bool
	for _, c := range cities {
		if c.Name == "TestCity" {
			found = true
			assert.Equal(t, len(commodities), len(c.MarketHall.Commodities))
			break
		}
	}
	assert.True(t, found, "Created city should be found in FindAll results")

	// Test Update
	city.Name = "UpdatedTestCity"
	err = repo.Update(&city)
	assert.NoError(t, err)

	// Test Delete
	err = repo.Delete(&city)
	assert.NoError(t, err)
}
