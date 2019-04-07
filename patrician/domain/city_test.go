package domain_test

import (
	"fmt"
	"testing"

	"github.com/adnvilla/patrician/patrician/domain"

	"github.com/stretchr/testify/assert"
)

func TestSupplyCommodity(t *testing.T) {

	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetCommodities(commodities)
	}

	city := domain.Cities["Estocolmo"]

	city.UpdateCommodity("Beer", 15.0, 16.0, 10, 11)

	stock := city.GetSupplyCommodityFromCity("Beer", "Visby")

	fmt.Println(stock)

	assert.NotNil(t, city)

}
