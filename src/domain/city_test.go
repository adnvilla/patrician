package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

// Helper function to prepare cities with commodities
func prepareCities() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}
}

// City commodity operations tests
func TestCityUpdateCommodity(t *testing.T) {
	// Setup
	commodities := domain.GetCommodities()
	city := &domain.City{
		Name: "TestCity",
		MarketHall: domain.MarketHall{
			Commodities: commodities,
		},
	}

	t.Run("UpdateExistingCommodity", func(t *testing.T) {
		err := city.UpdateCommodity("Beer", 15, 20, 50, 30)
		assert.NoError(t, err)

		beer := city.GetCommodities()["Beer"]
		assert.Equal(t, int16(15), beer.Buy)
		assert.Equal(t, int16(20), beer.Sell)
		assert.Equal(t, int16(50), beer.Production)
		assert.Equal(t, int16(30), beer.Consumption)
	})

	t.Run("UpdateNonExistentCommodity", func(t *testing.T) {
		err := city.UpdateCommodity("NonExistent", 10, 15, 20, 10)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "commodity NonExistent not found")
	})

	t.Run("UpdateWithNegativeValues", func(t *testing.T) {
		err := city.UpdateCommodity("Cloth", -5, -10, -15, -20)
		assert.NoError(t, err)

		cloth := city.GetCommodities()["Cloth"]
		assert.Equal(t, int16(-5), cloth.Buy)
		assert.Equal(t, int16(-10), cloth.Sell)
		assert.Equal(t, int16(-15), cloth.Production)
		assert.Equal(t, int16(-20), cloth.Consumption)
	})
}

func TestUpdateCommodityNotFound(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Unknown", 1, 1, 1, 1)
	assert.Error(t, err)
}

func TestCityGetCommodities(t *testing.T) {
	commodities := domain.GetCommodities()
	city := &domain.City{
		Name: "TestCity",
		MarketHall: domain.MarketHall{
			Commodities: commodities,
		},
	}

	result := city.GetCommodities()
	assert.Equal(t, len(commodities), len(result))
	assert.Equal(t, commodities, result)
}

func TestCitySetCommodities(t *testing.T) {
	city := &domain.City{Name: "TestCity"}
	newCommodities := map[string]*domain.Commodity{
		"TestCommodity": {Name: "TestCommodity"},
	}

	city.SetCommodities(newCommodities)
	assert.Equal(t, newCommodities, city.MarketHall.Commodities)
}

func TestCitySetMarketHall(t *testing.T) {
	city := &domain.City{Name: "TestCity"}
	marketHall := domain.MarketHall{
		Commodities: domain.GetCommodities(),
	}

	city.SetMarketHall(marketHall)
	assert.Equal(t, marketHall, city.MarketHall)
}

// Stock and supply tests
func TestCityGetStockCommodity(t *testing.T) {
	commodities := domain.GetCommodities()
	// Set specific production and consumption values
	commodities["Beer"].Production = 100
	commodities["Beer"].Consumption = 30

	city := &domain.City{
		Name: "TestCity",
		MarketHall: domain.MarketHall{
			Commodities: commodities,
		},
	}

	stock := city.GetStockCommodity("Beer")
	assert.Equal(t, int16(70), stock) // 100 - 30
}

func TestCityGetStockCommodities(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Beer", 15, 16, 10, 100)
	assert.NoError(t, err)

	stocks := city.GetStockCommodities()
	assert.Equal(t, 20, len(stocks))
	assert.Equal(t, int16(-90), stocks["Beer"])
}

func TestSupplyCommodity(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Beer", 15, 16, 10, 100)
	assert.NoError(t, err)
	stock := city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, int16(-30), stock)

	err = city.UpdateCommodity("Beer", 16, 17, 11, 110)
	assert.NoError(t, err)
	stock = city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, int16(-33), stock)

	assert.NotNil(t, city)
	assert.NotNil(t, city.MarketHall)
	assert.NotNil(t, city.MarketHall.Commodities)
	assert.Equal(t, 20, len(city.MarketHall.Commodities))
}

func TestCityGetSupplyCommodityFromCity(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Beer", 15, 16, 10, 100)
	assert.NoError(t, err)

	supply := city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, int16(-30), supply)
}

func TestCityGetSupplyCommoditiesFromCity(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Beer", 15, 16, 10, 100)
	assert.NoError(t, err)

	supply := city.GetSupplyCommoditiesFromCity("Visby")
	assert.Equal(t, 20, len(supply))
	assert.Equal(t, int16(-30), supply["Beer"])
}

// Default cities and entity tests
func TestDefaultCities(t *testing.T) {
	cities := domain.Cities

	assert.NotNil(t, cities)
	assert.Greater(t, len(cities), 0, "Cities map should not be empty")

	// Test that some expected cities exist
	expectedCities := []string{"Estocolmo", "Visby", "Londres", "Hamburgo"}
	for _, cityName := range expectedCities {
		city, exists := cities[cityName]
		assert.True(t, exists, "City %s should exist", cityName)
		assert.NotNil(t, city, "City %s should not be nil", cityName)
		assert.Equal(t, cityName, city.Name, "City name should match")
	}

	// Test that all cities have proper names
	for cityName, city := range cities {
		assert.Equal(t, cityName, city.Name, "City name should match key")
	}
}

func TestCityEntityProperties(t *testing.T) {
	city := &domain.City{Name: "TestCity"}

	// Test that the city has Entity properties
	assert.Equal(t, uint(0), city.ID)
	assert.True(t, city.CreatedAt.IsZero())

	// Test MarketHall initialization
	assert.NotNil(t, city.MarketHall)
}
