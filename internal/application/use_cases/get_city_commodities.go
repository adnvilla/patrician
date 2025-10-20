package use_cases

import (
	"context"

	"github.com/adnvilla/patrician/internal/domain"
	"github.com/adnvilla/patrician/internal/pkg/use_case"
)

type GetCityCommoditiesInput struct {
	CityName string
}

type GetCityCommoditiesUseCase struct {
}

func NewGetCityCommoditiesUseCase() use_case.UseCase[GetCityCommoditiesInput, map[string]*domain.Commodity] {
	return new(GetCityCommoditiesUseCase)
}

func (u *GetCityCommoditiesUseCase) Handle(ctx context.Context, in GetCityCommoditiesInput) (map[string]*domain.Commodity, error) {
	city := domain.Cities[in.CityName]
	return city.GetCommodities(), nil
}
