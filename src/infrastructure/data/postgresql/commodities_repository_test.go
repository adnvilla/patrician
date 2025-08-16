package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestCommodityRepositoryAdditional(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCommodityRepository(db)

	// Create test commodity
	commodity := domain.Commodity{
		Name:         "TestCommodity",
		Buy:          10,
		Sell:         15,
		Production:   100,
		Consumption:  80,
		CommodityType: domain.Barrel,
	}

	// Test Create
	err := repo.Create(&commodity)
	assert.NoError(t, err)
	assert.NotEqual(t, uint(0), commodity.ID)

	// Test FindByID
	retrieved, err := repo.FindByID(commodity.ID)
	assert.NoError(t, err)
	assert.Equal(t, "TestCommodity", retrieved.Name)
	assert.Equal(t, int16(10), retrieved.Buy)
	assert.Equal(t, int16(15), retrieved.Sell)
	assert.Equal(t, int16(100), retrieved.Production)
	assert.Equal(t, int16(80), retrieved.Consumption)
	assert.Equal(t, domain.Barrel, retrieved.CommodityType)

	// Test Update
	commodity.Buy = 12
	commodity.Sell = 18
	err = repo.Update(&commodity)
	assert.NoError(t, err)

	// Verify update
	updated, err := repo.FindByID(commodity.ID)
	assert.NoError(t, err)
	assert.Equal(t, int16(12), updated.Buy)
	assert.Equal(t, int16(18), updated.Sell)

	// Test Delete
	err = repo.Delete(&commodity)
	assert.NoError(t, err)

	// Verify deletion - should not be found
	_, err = repo.FindByID(commodity.ID)
	assert.Error(t, err)
}
