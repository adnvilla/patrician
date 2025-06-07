package use_cases_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCommoditiesUseCase(t *testing.T) {
	usecase := use_cases.NewGetCommoditiesUseCase()
	result, err := usecase.Handle(nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, domain.GetCommodities(), result)
}
