package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCommodityStockCalculations(t *testing.T) {
	testCases := []struct {
		name        string
		production  int16
		consumption int16
		expectedStock int16
	}{
		{"Zero production and consumption", 0, 0, 0},
		{"Equal production and consumption", 100, 100, 0},
		{"Higher production than consumption", 150, 100, 50},
		{"Higher consumption than production", 100, 150, -50},
		{"Maximum positive values", 32767, 0, 32767},
		{"Maximum negative values", 0, 32767, -32767},
		{"Negative production", -50, 100, -150},
		{"Negative consumption", 100, -50, 150},
		{"Both negative values", -100, -50, -50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			commodity := &domain.Commodity{
				Name:        "TestCommodity",
				Production:  tc.production,
				Consumption: tc.consumption,
			}

			stock := commodity.GetStock()
			assert.Equal(t, tc.expectedStock, stock)
		})
	}
}

func TestCommodityStructure(t *testing.T) {
	commodity := &domain.Commodity{
		Name:          "TestCommodity",
		CommodityType: domain.Barrel,
		Buy:           10,
		Sell:          15,
		Production:    100,
		Consumption:   50,
	}

	assert.Equal(t, "TestCommodity", commodity.Name)
	assert.Equal(t, domain.Barrel, commodity.CommodityType)
	assert.Equal(t, int16(10), commodity.Buy)
	assert.Equal(t, int16(15), commodity.Sell)
	assert.Equal(t, int16(100), commodity.Production)
	assert.Equal(t, int16(50), commodity.Consumption)
	assert.Equal(t, int16(50), commodity.GetStock())
}

func TestGetCommoditiesComprehensive(t *testing.T) {
	commodities := domain.GetCommodities()

	// Test that all expected commodities exist
	expectedCommodities := map[string]domain.CommodityType{
		"Beer":      domain.Barrel,
		"Bricks":    domain.Load,
		"Cloth":     domain.Barrel,
		"Fish":      domain.Load,
		"Grain":     domain.Load,
		"Hemp":      domain.Load,
		"Honey":     domain.Barrel,
		"IronGoods": domain.Barrel,
		"Leather":   domain.Barrel,
		"Meat":      domain.Load,
		"PigIron":   domain.Load,
		"Pitch":     domain.Barrel,
		"Pottery":   domain.Barrel,
		"Salt":      domain.Barrel,
		"Skins":     domain.Barrel,
		"Spices":    domain.Barrel,
		"Timber":    domain.Load,
		"WhaleOil":  domain.Barrel,
		"Wine":      domain.Barrel,
		"Wool":      domain.Load,
	}

	assert.Equal(t, len(expectedCommodities), len(commodities))

	for name, expectedType := range expectedCommodities {
		commodity, exists := commodities[name]
		assert.True(t, exists, "Commodity %s should exist", name)
		assert.Equal(t, name, commodity.Name)
		assert.Equal(t, expectedType, commodity.CommodityType)
	}
}

func TestCommodityTypeConsistency(t *testing.T) {
	commodities := domain.GetCommodities()

	barrelCommodities := []string{"Beer", "Cloth", "Honey", "IronGoods", "Leather", "Pitch", "Pottery", "Salt", "Skins", "Spices", "WhaleOil", "Wine"}
	loadCommodities := []string{"Bricks", "Fish", "Grain", "Hemp", "Meat", "PigIron", "Timber", "Wool"}

	for _, name := range barrelCommodities {
		commodity := commodities[name]
		assert.Equal(t, domain.Barrel, commodity.CommodityType, "Commodity %s should be of type Barrel", name)
	}

	for _, name := range loadCommodities {
		commodity := commodities[name]
		assert.Equal(t, domain.Load, commodity.CommodityType, "Commodity %s should be of type Load", name)
	}
}

func TestCommodityDefaultValues(t *testing.T) {
	commodities := domain.GetCommodities()

	for name, commodity := range commodities {
		// All default commodities should have zero values for buy, sell, production, consumption
		assert.Equal(t, int16(0), commodity.Buy, "Commodity %s should have zero Buy value", name)
		assert.Equal(t, int16(0), commodity.Sell, "Commodity %s should have zero Sell value", name)
		assert.Equal(t, int16(0), commodity.Production, "Commodity %s should have zero Production value", name)
		assert.Equal(t, int16(0), commodity.Consumption, "Commodity %s should have zero Consumption value", name)
		assert.Equal(t, int16(0), commodity.GetStock(), "Commodity %s should have zero Stock", name)
	}
}

func TestCommodityBoundaryValues(t *testing.T) {
	t.Run("MaximumPositiveValues", func(t *testing.T) {
		commodity := &domain.Commodity{
			Production:  32767, // Maximum int16 value
			Consumption: 0,
		}
		assert.Equal(t, int16(32767), commodity.GetStock())
	})

	t.Run("MaximumNegativeValues", func(t *testing.T) {
		commodity := &domain.Commodity{
			Production:  0,
			Consumption: 32767, // Maximum int16 value
		}
		assert.Equal(t, int16(-32767), commodity.GetStock())
	})

	t.Run("MinimumValues", func(t *testing.T) {
		commodity := &domain.Commodity{
			Production:  -32768, // Minimum int16 value
			Consumption: 0,
		}
		assert.Equal(t, int16(-32768), commodity.GetStock())
	})
}

func TestCommodityModification(t *testing.T) {
	commodities := domain.GetCommodities()
	
	// Modify a commodity
	beer := commodities["Beer"]
	originalBeer := *beer // Copy the original
	
	beer.Buy = 10
	beer.Sell = 15
	beer.Production = 100
	beer.Consumption = 30
	
	// Verify modifications
	assert.Equal(t, int16(10), beer.Buy)
	assert.Equal(t, int16(15), beer.Sell)
	assert.Equal(t, int16(100), beer.Production)
	assert.Equal(t, int16(30), beer.Consumption)
	assert.Equal(t, int16(70), beer.GetStock())
	
	// Verify name and type haven't changed
	assert.Equal(t, originalBeer.Name, beer.Name)
	assert.Equal(t, originalBeer.CommodityType, beer.CommodityType)
}
