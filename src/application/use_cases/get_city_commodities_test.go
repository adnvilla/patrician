package use_cases_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCityCommoditiesUseCase(t *testing.T) {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	usecase := use_cases.NewGetCityCommoditiesUseCase()
	input := use_cases.GetCityCommoditiesInput{CityName: "Estocolmo"}
	result, err := usecase.Handle(nil, input)
	assert.NoError(t, err)
	assert.Equal(t, domain.Cities["Estocolmo"].GetCommodities(), result)
}
