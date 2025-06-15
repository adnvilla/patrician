package main

import (
	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}

	router := gin.Default()

	router.Use(cors.Default())

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
