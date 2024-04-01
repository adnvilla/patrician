package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/patrician/domain"

	"github.com/stretchr/testify/assert"
)

func TestSupplyCommodity(t *testing.T) {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	city := domain.Cities["Estocolmo"]

	city.UpdateCommodity("Beer", 15, 16, 10, 100)
	stock := city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, int64(-30), stock)

	city.UpdateCommodity("Beer", 16, 17, 11, 110)
	stock = city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, int64(-33), stock)

	assert.NotNil(t, city)
	assert.NotNil(t, city.MarketHall)
	assert.NotNil(t, city.MarketHall.Commodities)
	assert.Equal(t, 20, len(city.MarketHall.Commodities))
}
