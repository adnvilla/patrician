package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestMarketHallStructure(t *testing.T) {
	marketHall := domain.MarketHall{
		Commodities: make(map[string]*domain.Commodity),
	}
	
	assert.NotNil(t, marketHall.Commodities)
	assert.Equal(t, 0, len(marketHall.Commodities))
	
	// Verify Entity inheritance
	assert.Equal(t, uint(0), marketHall.ID)
	assert.True(t, marketHall.CreatedAt.IsZero())
}

func TestMarketHallWithCommodities(t *testing.T) {
	commodities := domain.GetCommodities()
	marketHall := domain.MarketHall{
		Commodities: commodities,
	}
	
	assert.Equal(t, len(commodities), len(marketHall.Commodities))
	
	// Verify all commodities are properly stored
	for name, commodity := range commodities {
		storedCommodity, exists := marketHall.Commodities[name]
		assert.True(t, exists, "Commodity %s should exist in market hall", name)
		assert.Equal(t, commodity, storedCommodity)
	}
}

func TestMarketHallCommodityOperations(t *testing.T) {
	marketHall := domain.MarketHall{
		Commodities: make(map[string]*domain.Commodity),
	}
	
	// Add a commodity
	beer := &domain.Commodity{
		Name:          "Beer",
		CommodityType: domain.Barrel,
		Buy:           10,
		Sell:          15,
		Production:    100,
		Consumption:   50,
	}
	
	marketHall.Commodities["Beer"] = beer
	assert.Equal(t, 1, len(marketHall.Commodities))
	
	// Retrieve and verify commodity
	retrievedBeer, exists := marketHall.Commodities["Beer"]
	assert.True(t, exists)
	assert.Equal(t, beer, retrievedBeer)
	assert.Equal(t, "Beer", retrievedBeer.Name)
	assert.Equal(t, int16(50), retrievedBeer.GetStock())
}

func TestMarketHallEmpty(t *testing.T) {
	marketHall := domain.MarketHall{}
	
	// Test empty market hall
	assert.Nil(t, marketHall.Commodities)
	
	// Initialize commodities map
	marketHall.Commodities = make(map[string]*domain.Commodity)
	assert.NotNil(t, marketHall.Commodities)
	assert.Equal(t, 0, len(marketHall.Commodities))
}

func TestMarketHallMultipleCommodities(t *testing.T) {
	commodities := map[string]*domain.Commodity{
		"Beer": {
			Name:          "Beer",
			CommodityType: domain.Barrel,
			Production:    100,
			Consumption:   30,
		},
		"Cloth": {
			Name:          "Cloth",
			CommodityType: domain.Barrel,
			Production:    80,
			Consumption:   20,
		},
		"Fish": {
			Name:          "Fish",
			CommodityType: domain.Load,
			Production:    200,
			Consumption:   150,
		},
	}
	
	marketHall := domain.MarketHall{
		Commodities: commodities,
	}
	
	assert.Equal(t, 3, len(marketHall.Commodities))
	
	// Test each commodity
	assert.Equal(t, int16(70), marketHall.Commodities["Beer"].GetStock())
	assert.Equal(t, int16(60), marketHall.Commodities["Cloth"].GetStock())
	assert.Equal(t, int16(50), marketHall.Commodities["Fish"].GetStock())
}

func TestMarketHallCommodityModification(t *testing.T) {
	commodities := domain.GetCommodities()
	marketHall := domain.MarketHall{
		Commodities: commodities,
	}
	
	// Modify a commodity in the market hall
	beer := marketHall.Commodities["Beer"]
	beer.Production = 150
	beer.Consumption = 75
	
	// Verify the modification
	assert.Equal(t, int16(150), marketHall.Commodities["Beer"].Production)
	assert.Equal(t, int16(75), marketHall.Commodities["Beer"].Consumption)
	assert.Equal(t, int16(75), marketHall.Commodities["Beer"].GetStock())
}

func TestMarketHallNilCommodities(t *testing.T) {
	marketHall := domain.MarketHall{
		Commodities: nil,
	}
	
	assert.Nil(t, marketHall.Commodities)
	
	// Initialize with empty map
	marketHall.Commodities = make(map[string]*domain.Commodity)
	assert.NotNil(t, marketHall.Commodities)
	assert.Equal(t, 0, len(marketHall.Commodities))
}

func TestMarketHallEntityInheritance(t *testing.T) {
	marketHall := domain.MarketHall{
		Commodities: domain.GetCommodities(),
	}
	
	// Test Entity inheritance
	assert.Equal(t, uint(0), marketHall.Entity.ID)
	assert.True(t, marketHall.Entity.CreatedAt.IsZero())
	assert.True(t, marketHall.Entity.UpdatedAt.IsZero())
	assert.Nil(t, marketHall.Entity.DeletedAt.Time)
	
	// Test that commodities are properly stored
	assert.NotNil(t, marketHall.Commodities)
	assert.Equal(t, len(domain.GetCommodities()), len(marketHall.Commodities))
}

func TestMarketHallCommodityTypes(t *testing.T) {
	commodities := domain.GetCommodities()
	marketHall := domain.MarketHall{
		Commodities: commodities,
	}
	
	barrelCount := 0
	loadCount := 0
	
	for _, commodity := range marketHall.Commodities {
		switch commodity.CommodityType {
		case domain.Barrel:
			barrelCount++
		case domain.Load:
			loadCount++
		}
	}
	
	// Verify we have both types of commodities
	assert.Greater(t, barrelCount, 0, "Should have barrel commodities")
	assert.Greater(t, loadCount, 0, "Should have load commodities")
	assert.Equal(t, len(commodities), barrelCount+loadCount, "All commodities should be accounted for")
}
