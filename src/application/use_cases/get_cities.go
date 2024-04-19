package use_cases

import (
	"context"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/pkg/use_case"
)

type GetCitiesUseCase struct {
}

func NewGetCitiesUseCase() use_case.UseCase[any, map[string]*domain.City] {
	return new(GetCitiesUseCase)
}

func (u *GetCitiesUseCase) Handle(ctx context.Context, in any) (map[string]*domain.City, error) {
	return domain.Cities, nil
}
