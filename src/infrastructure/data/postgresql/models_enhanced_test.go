package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestCityModelStructure(t *testing.T) {
	cityModel := postgresql.CityModel{
		Name:         "TestCity",
		MarketHallID: 1,
	}

	// Test basic fields
	assert.Equal(t, "TestCity", cityModel.Name)
	assert.Equal(t, uint(1), cityModel.MarketHallID)

	// Test Entity inheritance
	assert.Equal(t, uint(0), cityModel.Entity.ID)
	assert.True(t, cityModel.Entity.CreatedAt.IsZero())
}

func TestCommodityModelStructure(t *testing.T) {
	commodityModel := postgresql.CommodityModel{
		Name:          "Beer",
		CommodityType: int(domain.Barrel),
		Buy:           10,
		Sell:          15,
		Production:    100,
		Consumption:   50,
		MarketHallID:  1,
	}

	// Test basic fields
	assert.Equal(t, "Beer", commodityModel.Name)
	assert.Equal(t, int(domain.Barrel), commodityModel.CommodityType)
	assert.Equal(t, int16(10), commodityModel.Buy)
	assert.Equal(t, int16(15), commodityModel.Sell)
	assert.Equal(t, int16(100), commodityModel.Production)
	assert.Equal(t, int16(50), commodityModel.Consumption)
	assert.Equal(t, uint(1), commodityModel.MarketHallID)

	// Test Entity inheritance
	assert.Equal(t, uint(0), commodityModel.Entity.ID)
}

func TestMarketHallModelStructure(t *testing.T) {
	commodities := []postgresql.CommodityModel{
		{
			Name:          "Beer",
			CommodityType: int(domain.Barrel),
			MarketHallID:  1,
		},
		{
			Name:          "Cloth",
			CommodityType: int(domain.Barrel),
			MarketHallID:  1,
		},
	}

	marketHallModel := postgresql.MarketHallModel{
		Commodities: commodities,
		CityID:      1,
	}

	// Test basic fields
	assert.Equal(t, uint(1), marketHallModel.CityID)
	assert.Equal(t, 2, len(marketHallModel.Commodities))

	// Test commodities
	assert.Equal(t, "Beer", marketHallModel.Commodities[0].Name)
	assert.Equal(t, "Cloth", marketHallModel.Commodities[1].Name)

	// Test Entity inheritance
	assert.Equal(t, uint(0), marketHallModel.Entity.ID)
}

func TestDistanceModelStructure(t *testing.T) {
	distanceModel := postgresql.DistanceModel{
		FromCity: "Estocolmo",
		ToCity:   "Visby",
		Value:    1.5,
	}

	// Test basic fields
	assert.Equal(t, "Estocolmo", distanceModel.FromCity)
	assert.Equal(t, "Visby", distanceModel.ToCity)
	assert.Equal(t, float32(1.5), distanceModel.Value)

	// Test Entity inheritance
	assert.Equal(t, uint(0), distanceModel.Entity.ID)
}

func TestCommodityTypeMapping(t *testing.T) {
	// Test mapping from domain types to int values
	barrelModel := postgresql.CommodityModel{
		CommodityType: int(domain.Barrel),
	}
	loadModel := postgresql.CommodityModel{
		CommodityType: int(domain.Load),
	}

	assert.Equal(t, int(domain.Barrel), barrelModel.CommodityType)
	assert.Equal(t, int(domain.Load), loadModel.CommodityType)
	assert.NotEqual(t, barrelModel.CommodityType, loadModel.CommodityType)
}

func TestModelRelationships(t *testing.T) {
	// Test the relationship structure - simulate what GORM would do
	marketHall := postgresql.MarketHallModel{
		CityID: 1,
	}
	// Simulate GORM setting the ID
	marketHall.Entity.ID = 1

	city := postgresql.CityModel{
		Name:         "TestCity",
		MarketHallID: 1,
		MarketHall:   marketHall,
	}

	commodity := postgresql.CommodityModel{
		Name:         "Beer",
		MarketHallID: 1,
	}

	// Test relationships
	assert.Equal(t, city.MarketHallID, marketHall.Entity.ID)      // Should match after GORM would set it
	assert.Equal(t, commodity.MarketHallID, marketHall.Entity.ID) // Should match after GORM would set it
}

func TestModelZeroValues(t *testing.T) {
	t.Run("CityModel", func(t *testing.T) {
		var city postgresql.CityModel
		assert.Equal(t, "", city.Name)
		assert.Equal(t, uint(0), city.MarketHallID)
		assert.Equal(t, uint(0), city.Entity.ID)
	})

	t.Run("CommodityModel", func(t *testing.T) {
		var commodity postgresql.CommodityModel
		assert.Equal(t, "", commodity.Name)
		assert.Equal(t, 0, commodity.CommodityType)
		assert.Equal(t, int16(0), commodity.Buy)
		assert.Equal(t, int16(0), commodity.Sell)
		assert.Equal(t, int16(0), commodity.Production)
		assert.Equal(t, int16(0), commodity.Consumption)
		assert.Equal(t, uint(0), commodity.MarketHallID)
	})

	t.Run("MarketHallModel", func(t *testing.T) {
		var marketHall postgresql.MarketHallModel
		assert.Nil(t, marketHall.Commodities)
		assert.Equal(t, uint(0), marketHall.CityID)
	})

	t.Run("DistanceModel", func(t *testing.T) {
		var distance postgresql.DistanceModel
		assert.Equal(t, "", distance.FromCity)
		assert.Equal(t, "", distance.ToCity)
		assert.Equal(t, float32(0), distance.Value)
	})
}

func TestModelFieldTypes(t *testing.T) {
	t.Run("CityModel", func(t *testing.T) {
		city := postgresql.CityModel{}
		assert.IsType(t, "", city.Name)
		assert.IsType(t, uint(0), city.MarketHallID)
	})

	t.Run("CommodityModel", func(t *testing.T) {
		commodity := postgresql.CommodityModel{}
		assert.IsType(t, "", commodity.Name)
		assert.IsType(t, 0, commodity.CommodityType)
		assert.IsType(t, int16(0), commodity.Buy)
		assert.IsType(t, int16(0), commodity.Sell)
		assert.IsType(t, int16(0), commodity.Production)
		assert.IsType(t, int16(0), commodity.Consumption)
		assert.IsType(t, uint(0), commodity.MarketHallID)
	})

	t.Run("DistanceModel", func(t *testing.T) {
		distance := postgresql.DistanceModel{}
		assert.IsType(t, "", distance.FromCity)
		assert.IsType(t, "", distance.ToCity)
		assert.IsType(t, float32(0), distance.Value)
	})
}

func TestModelInstantiation(t *testing.T) {
	// Test creating models with realistic data
	city := postgresql.CityModel{
		Name:         "Estocolmo",
		MarketHallID: 1,
	}
	assert.NotNil(t, city)
	assert.Equal(t, "Estocolmo", city.Name)

	commodity := postgresql.CommodityModel{
		Name:          "Beer",
		CommodityType: int(domain.Barrel),
		Buy:           15,
		Sell:          20,
		Production:    100,
		Consumption:   30,
		MarketHallID:  1,
	}
	assert.NotNil(t, commodity)
	assert.Equal(t, "Beer", commodity.Name)

	marketHall := postgresql.MarketHallModel{
		Commodities: []postgresql.CommodityModel{commodity},
		CityID:      1,
	}
	assert.NotNil(t, marketHall)
	assert.Equal(t, 1, len(marketHall.Commodities))

	distance := postgresql.DistanceModel{
		FromCity: "Estocolmo",
		ToCity:   "Visby",
		Value:    1.5,
	}
	assert.NotNil(t, distance)
	assert.Equal(t, "Estocolmo", distance.FromCity)
}
