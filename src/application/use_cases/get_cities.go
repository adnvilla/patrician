package use_cases

import (
	"context"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/pkg/use_case"
)

type GetCitiesUseCase struct {
}

func NewGetCitiesUseCase() use_case.UseCase {
	return new(GetCitiesUseCase)
}

func (u *GetCitiesUseCase) Handle(ctx context.Context, in use_case.Input) (use_case.Output, error) {
	return domain.Cities, nil
}
