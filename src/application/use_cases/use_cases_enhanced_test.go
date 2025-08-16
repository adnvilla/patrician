package use_cases_test

import (
	"context"
	"testing"

	"github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCitiesUseCaseEnhanced(t *testing.T) {
	usecase := use_cases.NewGetCitiesUseCase()

	t.Run("SuccessfulExecution", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(domain.Cities), len(result))

		// Verify all cities are present
		for cityName, expectedCity := range domain.Cities {
			actualCity, exists := result[cityName]
			assert.True(t, exists, "City %s should exist in result", cityName)
			assert.Equal(t, expectedCity.Name, actualCity.Name)
		}
	})

	t.Run("WithContext", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "test", "value")
		result, err := usecase.Handle(ctx, nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(domain.Cities), len(result))
	})

	t.Run("WithNilInput", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("MultipleExecutions", func(t *testing.T) {
		// Test that multiple executions return consistent results
		result1, err1 := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err1)

		result2, err2 := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err2)

		assert.Equal(t, len(result1), len(result2))
	})
}

func TestGetCommoditiesUseCaseEnhanced(t *testing.T) {
	usecase := use_cases.NewGetCommoditiesUseCase()

	t.Run("SuccessfulExecution", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(domain.GetCommodities()), len(result))

		// Verify all commodities are present
		expectedCommodities := domain.GetCommodities()
		for commodityName, expectedCommodity := range expectedCommodities {
			actualCommodity, exists := result[commodityName]
			assert.True(t, exists, "Commodity %s should exist in result", commodityName)
			assert.Equal(t, expectedCommodity.Name, actualCommodity.Name)
			assert.Equal(t, expectedCommodity.CommodityType, actualCommodity.CommodityType)
		}
	})

	t.Run("CommodityTypes", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)

		barrelCount := 0
		loadCount := 0

		for _, commodity := range result {
			switch commodity.CommodityType {
			case domain.Barrel:
				barrelCount++
			case domain.Load:
				loadCount++
			}
		}

		assert.Greater(t, barrelCount, 0, "Should have barrel commodities")
		assert.Greater(t, loadCount, 0, "Should have load commodities")
	})

	t.Run("WithDifferentInputs", func(t *testing.T) {
		// Should work with any input since it's ignored
		inputs := []interface{}{nil, "test", 123, map[string]string{"key": "value"}}

		for _, input := range inputs {
			result, err := usecase.Handle(context.Background(), input)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, len(domain.GetCommodities()), len(result))
		}
	})
}

func TestGetDistancesUseCaseEnhanced(t *testing.T) {
	usecase := use_cases.NewGetDistancesUseCase()

	t.Run("SuccessfulExecution", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(domain.Distances), len(result))

		// Verify all cities have distance entries
		for cityName := range domain.Cities {
			distances, exists := result[cityName]
			assert.True(t, exists, "City %s should have distance entries", cityName)
			assert.NotNil(t, distances)
		}
	})

	t.Run("DistanceSymmetry", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)

		// Test that distances are symmetric
		for fromCity, fromDistances := range result {
			for toCity, distance := range fromDistances {
				if toDistances, exists := result[toCity]; exists {
					if reverseDistance, exists := toDistances[fromCity]; exists {
						assert.Equal(t, distance, reverseDistance,
							"Distance should be symmetric between %s and %s", fromCity, toCity)
					}
				}
			}
		}
	})

	t.Run("PositiveDistances", func(t *testing.T) {
		result, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)

		// All distances should be positive
		for fromCity, fromDistances := range result {
			for toCity, distance := range fromDistances {
				assert.Greater(t, distance, float32(0),
					"Distance from %s to %s should be positive", fromCity, toCity)
			}
		}
	})
}

func TestGetCityCommoditiesUseCaseEnhanced(t *testing.T) {
	prepareCities()
	usecase := use_cases.NewGetCityCommoditiesUseCase()

	t.Run("ValidCity", func(t *testing.T) {
		input := use_cases.GetCityCommoditiesInput{CityName: "Estocolmo"}
		result, err := usecase.Handle(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(domain.GetCommodities()), len(result))
	})

	t.Run("AllCities", func(t *testing.T) {
		for cityName := range domain.Cities {
			input := use_cases.GetCityCommoditiesInput{CityName: cityName}
			result, err := usecase.Handle(context.Background(), input)
			assert.NoError(t, err, "Should handle city %s without error", cityName)
			assert.NotNil(t, result, "Result should not be nil for city %s", cityName)
		}
	})

	t.Run("EmptyInput", func(t *testing.T) {
		input := use_cases.GetCityCommoditiesInput{CityName: ""}
		result, err := usecase.Handle(context.Background(), input)
		// This might return an error or empty result depending on implementation
		// For now, we just verify it doesn't panic
		_ = result
		_ = err
	})

	t.Run("NonExistentCity", func(t *testing.T) {
		input := use_cases.GetCityCommoditiesInput{CityName: "NonExistentCity"}
		result, err := usecase.Handle(context.Background(), input)
		// This might return an error or empty result depending on implementation
		// For now, we just verify it doesn't panic
		_ = result
		_ = err
	})

	t.Run("CommodityStructure", func(t *testing.T) {
		input := use_cases.GetCityCommoditiesInput{CityName: "Estocolmo"}
		result, err := usecase.Handle(context.Background(), input)
		assert.NoError(t, err)

		// Verify that each commodity has the expected structure
		for commodityName, commodity := range result {
			assert.Equal(t, commodityName, commodity.Name)
			assert.Contains(t, []domain.CommodityType{domain.Barrel, domain.Load}, commodity.CommodityType)
		}
	})
}

func TestUseCaseCreation(t *testing.T) {
	t.Run("GetCitiesUseCase", func(t *testing.T) {
		usecase := use_cases.NewGetCitiesUseCase()
		assert.NotNil(t, usecase)

		// Test that it implements the UseCase interface
		_, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
	})

	t.Run("GetCommoditiesUseCase", func(t *testing.T) {
		usecase := use_cases.NewGetCommoditiesUseCase()
		assert.NotNil(t, usecase)

		_, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
	})

	t.Run("GetDistancesUseCase", func(t *testing.T) {
		usecase := use_cases.NewGetDistancesUseCase()
		assert.NotNil(t, usecase)

		_, err := usecase.Handle(context.Background(), nil)
		assert.NoError(t, err)
	})

	t.Run("GetCityCommoditiesUseCase", func(t *testing.T) {
		usecase := use_cases.NewGetCityCommoditiesUseCase()
		assert.NotNil(t, usecase)

		// Need to prepare cities first
		prepareCities()
		input := use_cases.GetCityCommoditiesInput{CityName: "Estocolmo"}
		_, err := usecase.Handle(context.Background(), input)
		assert.NoError(t, err)
	})
}
