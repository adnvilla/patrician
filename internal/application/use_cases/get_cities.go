package use_cases

import (
	"context"

	"github.com/adnvilla/patrician/internal/domain"
	"github.com/adnvilla/patrician/internal/pkg/use_case"
)

type GetCitiesUseCase struct {
}

func NewGetCitiesUseCase() use_case.UseCase[any, map[string]*domain.City] {
	return new(GetCitiesUseCase)
}

func (u *GetCitiesUseCase) Handle(ctx context.Context, in any) (map[string]*domain.City, error) {
	return domain.Cities, nil
}
