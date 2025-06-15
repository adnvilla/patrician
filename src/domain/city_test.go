package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"

	"github.com/stretchr/testify/assert"
)

func TestSupplyCommodity(t *testing.T) {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

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

func TestUpdateCommodityNotFound(t *testing.T) {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Unknown", 1, 1, 1, 1)
	assert.Error(t, err)
}
