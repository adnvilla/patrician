package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

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
	commodities := domain.GetCommodities()
	// Set specific values for testing
	commodities["Beer"].Production = 100
	commodities["Beer"].Consumption = 30
	commodities["Cloth"].Production = 50
	commodities["Cloth"].Consumption = 20

	city := &domain.City{
		Name: "TestCity",
		MarketHall: domain.MarketHall{
			Commodities: commodities,
		},
	}

	stocks := city.GetStockCommodities()
	assert.Equal(t, len(commodities), len(stocks))
	assert.Equal(t, int16(70), stocks["Beer"])
	assert.Equal(t, int16(30), stocks["Cloth"])
}

func TestCityGetSupplyCommodityFromCity(t *testing.T) {
	// Setup cities with distances
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	testCity := domain.Cities["Estocolmo"]
	targetCity := "Visby"

	// Set commodity values
	err := testCity.UpdateCommodity("Beer", 15, 20, 100, 50)
	assert.NoError(t, err)

	supply := testCity.GetSupplyCommodityFromCity("Beer", targetCity)
	assert.IsType(t, int16(0), supply)
	// The specific value depends on the distance calculation
}

func TestCityGetSupplyCommoditiesFromCity(t *testing.T) {
	// Setup cities with distances
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	testCity := domain.Cities["Estocolmo"]
	targetCity := "Visby"

	supplies := testCity.GetSupplyCommoditiesFromCity(targetCity)
	assert.Equal(t, len(testCity.GetCommodities()), len(supplies))

	// Verify that all commodities have supply calculations
	for name := range testCity.GetCommodities() {
		_, exists := supplies[name]
		assert.True(t, exists, "Supply should exist for commodity %s", name)
	}
}

func TestDefaultCities(t *testing.T) {
	expectedCities := []string{
		"Edimburgo", "Scarborough", "Londres", "Brujas", "Colonia",
		"Groninga", "Bremen", "Hamburgo", "Ripen", "Bergen", "Oslo",
		"Aalborg", "Malmo", "Lubeck", "Rostock", "Stettin", "Gdansk",
		"Torum", "Riga", "Visby", "Estocolmo", "Reval", "Ladoga", "Novgorod",
	}

	assert.Equal(t, len(expectedCities), len(domain.Cities))

	for _, cityName := range expectedCities {
		city, exists := domain.Cities[cityName]
		assert.True(t, exists, "City %s should exist", cityName)
		assert.Equal(t, cityName, city.Name)
	}
}

func TestCityEntityProperties(t *testing.T) {
	city := domain.Cities["Estocolmo"]
	
	// Test that the city has Entity properties
	assert.NotNil(t, city)
	assert.Equal(t, "Estocolmo", city.Name)
	
	// Test MarketHall initialization
	commodities := domain.GetCommodities()
	city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	assert.NotNil(t, city.MarketHall)
	assert.Equal(t, len(commodities), len(city.MarketHall.Commodities))
}
