package use_cases_test

import (
	"context"
	"testing"

	"github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func prepareCities() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}
}

func TestGetCitiesUseCase(t *testing.T) {
	usecase := use_cases.NewGetCitiesUseCase()
	result, err := usecase.Handle(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, len(domain.Cities), len(result))
}

func TestGetCityCommoditiesUseCase(t *testing.T) {
	prepareCities()
	usecase := use_cases.NewGetCityCommoditiesUseCase()
	input := use_cases.GetCityCommoditiesInput{CityName: "Estocolmo"}
	result, err := usecase.Handle(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, len(domain.GetCommodities()), len(result))
}

func TestGetCommoditiesUseCase(t *testing.T) {
	usecase := use_cases.NewGetCommoditiesUseCase()
	result, err := usecase.Handle(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, len(domain.GetCommodities()), len(result))
}

func TestGetDistancesUseCase(t *testing.T) {
	usecase := use_cases.NewGetDistancesUseCase()
	result, err := usecase.Handle(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, len(domain.Distances), len(result))
}
