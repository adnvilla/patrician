package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestMarketHallRepositoryAdditional(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewMarketHallRepository(db)

	// Create test market hall with commodities
	commodities := domain.GetCommodities()
	marketHall := domain.MarketHall{
		Commodities: commodities,
	}

	// Test Create
	err := repo.Create(&marketHall)
	assert.NoError(t, err)
	assert.NotEqual(t, uint(0), marketHall.ID)

	// Test Update - modify a commodity in the market hall
	marketHall.Commodities["Beer"].Buy = 12
	marketHall.Commodities["Beer"].Sell = 18
	err = repo.Update(&marketHall)
	assert.NoError(t, err)

	// Verify the update worked by retrieving the market hall
	retrieved, err := repo.FindByID(marketHall.ID)
	assert.NoError(t, err)
	assert.Equal(t, int16(12), retrieved.Commodities["Beer"].Buy)
	assert.Equal(t, int16(18), retrieved.Commodities["Beer"].Sell)

	// Test that all commodities are present
	assert.Equal(t, len(domain.GetCommodities()), len(retrieved.Commodities))

	// Verify specific commodities exist
	for name := range domain.GetCommodities() {
		commodity, exists := retrieved.Commodities[name]
		assert.True(t, exists, "Commodity %s should exist in market hall", name)
		assert.Equal(t, name, commodity.Name)
	}

	// Test Delete
	err = repo.Delete(&marketHall)
	assert.NoError(t, err)
}
