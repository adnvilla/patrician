package domain_test

import (
	"testing"
	"time"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestEntityStructure(t *testing.T) {
	entity := domain.Entity{}
	
	// Test that Entity has gorm.Model embedded
	assert.Equal(t, uint(0), entity.ID)
	assert.True(t, entity.CreatedAt.IsZero())
	assert.True(t, entity.UpdatedAt.IsZero())
	assert.Nil(t, entity.DeletedAt.Time)
}

func TestEntityInheritance(t *testing.T) {
	// Test that City inherits from Entity
	city := domain.City{
		Name: "TestCity",
	}
	
	// City should have Entity fields
	assert.Equal(t, uint(0), city.ID)
	assert.True(t, city.CreatedAt.IsZero())
	assert.True(t, city.UpdatedAt.IsZero())
	assert.Nil(t, city.DeletedAt.Time)
	assert.Equal(t, "TestCity", city.Name)
	
	// Test that Commodity inherits from Entity
	commodity := domain.Commodity{
		Name: "TestCommodity",
		CommodityType: domain.Barrel,
	}
	
	// Commodity should have Entity fields
	assert.Equal(t, uint(0), commodity.ID)
	assert.True(t, commodity.CreatedAt.IsZero())
	assert.True(t, commodity.UpdatedAt.IsZero())
	assert.Nil(t, commodity.DeletedAt.Time)
	assert.Equal(t, "TestCommodity", commodity.Name)
	assert.Equal(t, domain.Barrel, commodity.CommodityType)
	
	// Test that MarketHall inherits from Entity
	marketHall := domain.MarketHall{
		Commodities: make(map[string]*domain.Commodity),
	}
	
	// MarketHall should have Entity fields
	assert.Equal(t, uint(0), marketHall.ID)
	assert.True(t, marketHall.CreatedAt.IsZero())
	assert.True(t, marketHall.UpdatedAt.IsZero())
	assert.Nil(t, marketHall.DeletedAt.Time)
	assert.NotNil(t, marketHall.Commodities)
}

func TestEntityFieldTypes(t *testing.T) {
	entity := domain.Entity{}
	
	// Test field types
	assert.IsType(t, uint(0), entity.ID)
	assert.IsType(t, time.Time{}, entity.CreatedAt)
	assert.IsType(t, time.Time{}, entity.UpdatedAt)
	assert.IsType(t, (*time.Time)(nil), entity.DeletedAt.Time)
}

func TestEntityZeroValues(t *testing.T) {
	entity := domain.Entity{}
	
	// Test zero values
	assert.Zero(t, entity.ID)
	assert.Zero(t, entity.CreatedAt)
	assert.Zero(t, entity.UpdatedAt)
	assert.Nil(t, entity.DeletedAt.Time)
}

func TestEntityCityComposition(t *testing.T) {
	city := &domain.City{
		Name: "TestCity",
		MarketHall: domain.MarketHall{
			Commodities: domain.GetCommodities(),
		},
	}
	
	// Verify that City properly composes Entity
	assert.Equal(t, "TestCity", city.Name)
	assert.NotNil(t, city.MarketHall)
	assert.Equal(t, len(domain.GetCommodities()), len(city.MarketHall.Commodities))
	
	// Verify Entity fields are accessible
	assert.Equal(t, uint(0), city.Entity.ID)
	assert.True(t, city.Entity.CreatedAt.IsZero())
}

func TestEntityCommodityComposition(t *testing.T) {
	commodity := &domain.Commodity{
		Name:          "TestCommodity",
		CommodityType: domain.Load,
		Buy:           10,
		Sell:          15,
		Production:    100,
		Consumption:   50,
	}
	
	// Verify that Commodity properly composes Entity
	assert.Equal(t, "TestCommodity", commodity.Name)
	assert.Equal(t, domain.Load, commodity.CommodityType)
	assert.Equal(t, int16(10), commodity.Buy)
	assert.Equal(t, int16(15), commodity.Sell)
	assert.Equal(t, int16(100), commodity.Production)
	assert.Equal(t, int16(50), commodity.Consumption)
	
	// Verify Entity fields are accessible
	assert.Equal(t, uint(0), commodity.Entity.ID)
	assert.True(t, commodity.Entity.CreatedAt.IsZero())
}
