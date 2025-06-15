package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func prepareCities() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}
}

func TestGetStockCommodities(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Beer", 15, 16, 10, 100)
	assert.NoError(t, err)

	stocks := city.GetStockCommodities()
	assert.Equal(t, 20, len(stocks))
	assert.Equal(t, int16(-90), stocks["Beer"])
}

func TestGetSupplyCommoditiesFromCity(t *testing.T) {
	prepareCities()
	city := domain.Cities["Estocolmo"]

	err := city.UpdateCommodity("Beer", 15, 16, 10, 100)
	assert.NoError(t, err)

	supply := city.GetSupplyCommoditiesFromCity("Visby")
	assert.Equal(t, 20, len(supply))
	assert.Equal(t, int16(-30), supply["Beer"])
}
