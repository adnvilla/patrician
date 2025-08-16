package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCommodityGetStockExtended(t *testing.T) {
	// Create test cases with different production and consumption values
	testCases := []struct {
		name        string
		production  int16
		consumption int16
		expected    int16
	}{
		{"Production equals consumption", 100, 100, 0},
		{"Production exceeds consumption", 150, 100, 50},
		{"Consumption exceeds production", 100, 150, -50},
		{"Zero production and consumption", 0, 0, 0},
		{"Only production, no consumption", 100, 0, 100},
		{"Only consumption, no production", 0, 100, -100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			commodity := domain.Commodity{
				Name:        "TestCommodity",
				Production:  tc.production,
				Consumption: tc.consumption,
			}

			// Test the GetStock method
			stock := commodity.GetStock()
			assert.Equal(t, tc.expected, stock)
		})
	}
}

func TestGetCommoditiesAdditional(t *testing.T) {
	// Get the commodities
	commodities := domain.GetCommodities()

	// Check for specific commodities that should exist
	expectedCommodities := []string{
		"Beer", "Bricks", "Cloth", "Fish", "Grain",
		"Hemp", "Honey", "IronGoods", "Leather", "Meat",
		"PigIron", "Pitch", "Pottery", "Salt", "Skins",
		"Spices", "Timber", "WhaleOil", "Wine", "Wool",
	}

	for _, name := range expectedCommodities {
		commodity, exists := commodities[name]
		assert.True(t, exists, "Commodity %s should exist", name)
		assert.Equal(t, name, commodity.Name)
	}
}

func TestCommodityTypes(t *testing.T) {
	commodities := domain.GetCommodities()

	// Check that each commodity has the expected CommodityType
	barrelCommodities := []string{
		"Beer", "Cloth", "Honey", "IronGoods", "Leather",
		"Pitch", "Pottery", "Salt", "Skins", "Spices",
		"WhaleOil", "Wine",
	}

	loadCommodities := []string{
		"Bricks", "Fish", "Grain", "Hemp", "Meat",
		"PigIron", "Timber", "Wool",
	}

	for _, name := range barrelCommodities {
		assert.Equal(t, domain.Barrel, commodities[name].CommodityType,
			"Commodity %s should be of type Barrel", name)
	}

	for _, name := range loadCommodities {
		assert.Equal(t, domain.Load, commodities[name].CommodityType,
			"Commodity %s should be of type Load", name)
	}
}
