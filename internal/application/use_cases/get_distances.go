package use_cases

import (
	"context"

	"github.com/adnvilla/patrician/internal/domain"
	"github.com/adnvilla/patrician/internal/pkg/use_case"
)

type GetDistancesUseCase struct {
}

func NewGetDistancesUseCase() use_case.UseCase[any, map[string]map[string]float32] {
	return new(GetDistancesUseCase)
}

func (u *GetDistancesUseCase) Handle(ctx context.Context, in any) (map[string]map[string]float32, error) {
	return domain.Distances, nil
}
