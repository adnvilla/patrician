package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestMarketHall(t *testing.T) {
	// Create a market hall
	marketHall := domain.MarketHall{
		Commodities: domain.GetCommodities(),
	}

	// Test that the market hall has all commodities
	assert.Equal(t, len(domain.GetCommodities()), len(marketHall.Commodities))

	// Test accessing specific commodities
	beer, exists := marketHall.Commodities["Beer"]
	assert.True(t, exists, "Beer should exist in market hall")
	assert.Equal(t, "Beer", beer.Name)

	// Test that the commodities map is properly populated
	for name, commodity := range marketHall.Commodities {
		assert.Equal(t, name, commodity.Name)
	}

	// Test modifying a commodity in the market hall
	marketHall.Commodities["Beer"].Buy = 10
	marketHall.Commodities["Beer"].Sell = 20
	marketHall.Commodities["Beer"].Production = 100
	marketHall.Commodities["Beer"].Consumption = 80

	// Verify the modifications
	assert.Equal(t, int16(10), marketHall.Commodities["Beer"].Buy)
	assert.Equal(t, int16(20), marketHall.Commodities["Beer"].Sell)
	assert.Equal(t, int16(100), marketHall.Commodities["Beer"].Production)
	assert.Equal(t, int16(80), marketHall.Commodities["Beer"].Consumption)

	// Test the stock calculation
	stock := marketHall.Commodities["Beer"].GetStock()
	assert.Equal(t, int16(20), stock) // 100 - 80 = 20
}
