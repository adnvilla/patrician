package main

import (
	"context"
	"log"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	router := gin.Default()

	router.Use(cors.Default())

	swagger, err := openapi3.NewLoader().LoadFromFile("docs/openapi.yaml")
	if err != nil {
		log.Fatalf("load swagger: %v", err)
	}
	if err = swagger.Validate(context.Background()); err != nil {
		log.Fatalf("validate swagger: %v", err)
	}

	router.Use(middleware.OapiRequestValidator(swagger))

	router.GET("/cities", handlers.GetCities)
	router.GET("/commodities", handlers.GetCommodities)
	router.GET("/distances", handlers.GetDistances)
	router.GET("/city/:name/commodities", handlers.GetCityCommodities)
	router.POST("/city/:name/commodity", handlers.UpdateCommodity)
	router.POST("/city/:name/commodities", handlers.UpdateCommodities)
	router.GET("/city/:name/stock", handlers.GetStock)
	router.GET("/city/:name/supply/:city", handlers.GetSupply)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
