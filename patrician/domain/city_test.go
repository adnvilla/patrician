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

	city.UpdateCommodity("Beer", 15.0, 16.0, 10, 11)
	stock := city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, float64(-0.011428571428571429), stock)

	city.UpdateCommodity("Beer", 16.0, 17.0, 11, 13)
	stock = city.GetSupplyCommodityFromCity("Beer", "Visby")
	assert.Equal(t, float64(-0.022857142857142857), stock)

	assert.NotNil(t, city)
	assert.NotNil(t, city.MarketHall)
	assert.NotNil(t, city.MarketHall.Commodities)
	assert.Equal(t, 20, len(city.MarketHall.Commodities))
}
