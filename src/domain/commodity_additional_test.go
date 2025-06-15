package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCommodityGetStock(t *testing.T) {
	c := domain.Commodity{Production: 10, Consumption: 7}
	assert.Equal(t, int16(3), c.GetStock())
}

func TestGetCommodities(t *testing.T) {
	commodities := domain.GetCommodities()
	assert.Equal(t, 20, len(commodities))
	_, ok := commodities["Beer"]
	assert.True(t, ok)
}
