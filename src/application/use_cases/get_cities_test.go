package use_cases_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCitiesUseCase(t *testing.T) {
	usecase := use_cases.NewGetCitiesUseCase()
	result, err := usecase.Handle(nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, domain.Cities, result)
}
