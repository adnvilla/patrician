package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	// Initialize domain data
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Load OpenAPI spec
	swagger, err := openapi3.NewLoader().LoadFromFile("docs/openapi.yaml")
	if err != nil {
		panic(err)
	}
	if err = swagger.Validate(context.Background()); err != nil {
		panic(err)
	}
	router.Use(middleware.OapiRequestValidator(swagger))

	// Register routes
	router.GET("/cities", handlers.GetCities)
	router.GET("/commodities", handlers.GetCommodities)
	router.GET("/distances", handlers.GetDistances)
	router.GET("/city/:name/commodities", handlers.GetCityCommodities)
	router.POST("/city/:name/commodity", handlers.UpdateCommodity)
	router.POST("/city/:name/commodities", handlers.UpdateCommodities)
	router.GET("/city/:name/stock", handlers.GetStock)
	router.GET("/city/:name/supply/:city", handlers.GetSupply)

	return router
}

func TestMainRoutesHealthCheck(t *testing.T) {
	router := setupTestRouter()

	// Test main routes to ensure they're all working
	routes := []string{
		"/cities",
		"/commodities",
		"/distances",
		"/city/Estocolmo/commodities",
		"/city/Estocolmo/stock",
		"/city/Estocolmo/supply/Visby",
	}

	for _, route := range routes {
		req, _ := http.NewRequest(http.MethodGet, route, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Route %s should return 200", route)
	}
}

func TestCityDataConsistency(t *testing.T) {
	// Initialize domain data
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	// Check that all cities have the proper number of commodities
	for name, city := range domain.Cities {
		assert.NotNil(t, city.MarketHall, "City %s should have a market hall", name)
		assert.Equal(t, len(domain.GetCommodities()), len(city.MarketHall.Commodities),
			"City %s should have all commodities", name)
	}

	// Check that all distances are properly defined
	for fromCity, distancesMap := range domain.Distances {
		for toCity, _ := range distancesMap {
			// Check that both cities exist in the Cities map
			fromCityObj, fromExists := domain.Cities[fromCity]
			assert.True(t, fromExists, "City %s in distance data should exist in Cities map", fromCity)
			assert.NotNil(t, fromCityObj)

			toCityObj, toExists := domain.Cities[toCity]
			assert.True(t, toExists, "City %s in distance data should exist in Cities map", toCity)
			assert.NotNil(t, toCityObj)
		}
	}
}
