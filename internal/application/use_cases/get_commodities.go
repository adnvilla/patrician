package use_cases

import (
	"context"

	"github.com/adnvilla/patrician/internal/domain"
	"github.com/adnvilla/patrician/internal/pkg/use_case"
)

type GetCommoditiesUseCase struct {
}

func NewGetCommoditiesUseCase() use_case.UseCase[any, map[string]*domain.Commodity] {
	return new(GetCommoditiesUseCase)
}

func (u *GetCommoditiesUseCase) Handle(ctx context.Context, in any) (map[string]*domain.Commodity, error) {
	return domain.GetCommodities(), nil
}
