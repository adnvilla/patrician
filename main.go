package main

import (
	"net/http"

	"github.com/adnvilla/patrician/patrician/domain"
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

	router.GET("/cities", getCities)
	router.GET("/commodities", getCommodities)
	router.GET("/distances", getDistances)
	router.GET("/city/:name/commodities", getCityCommodities)
	router.POST("/city/:name/commodity", updateCommodity)
	router.POST("/city/:name/commodities", updateCommodities)
	router.GET("/city/:name/stock", getStock)
	router.GET("/city/:name/supply/:city", getSupply)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}

func getCities(c *gin.Context) {
	c.JSON(200, domain.Cities)
}

func getCommodities(c *gin.Context) {
	c.JSON(200, domain.GetCommodities())
}

func getDistances(c *gin.Context) {
	c.JSON(200, domain.Distances)
}

func getCityCommodities(c *gin.Context) {
	name := c.Param("name")

	city := domain.Cities[name]

	c.JSON(200, city.GetCommodities())
}

func getStock(c *gin.Context) {
	name := c.Param("name")

	city := domain.Cities[name]

	supply := city.GetStockCommodities()

	c.JSON(200, supply)
}

func getSupply(c *gin.Context) {
	name := c.Param("name")
	from := c.Param("city")

	city := domain.Cities[name]

	supply := city.GetSupplyCommoditiesFromCity(from)

	c.JSON(200, supply)
}

func updateCommodity(c *gin.Context) {
	name := c.Param("name")

	var commodity Commodity

	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := domain.Cities[name]

	city.UpdateCommodity(commodity.Name, commodity.Buy, commodity.Sell, commodity.Production, commodity.Consumption)

	c.String(http.StatusOK, "updated!")
}

func updateCommodities(c *gin.Context) {
	name := c.Param("name")

	var commodities Commodities

	if err := c.ShouldBindJSON(&commodities); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := domain.Cities[name]

	for _, commodity := range commodities.Commodities {
		city.UpdateCommodity(commodity.Name, commodity.Buy, commodity.Sell, commodity.Production, commodity.Consumption)
	}

	c.String(http.StatusOK, "updated!")
}

type Commodity struct {
	Name        string `json:"name" binding:"required"`
	Buy         int64  `json:"buy" binding:"required"`
	Sell        int64  `json:"sell" binding:"required"`
	Production  int64  `json:"production" binding:"required"`
	Consumption int64  `json:"consumption" binding:"required"`
}

type Commodities struct {
	Commodities []Commodity `json:"commodities" binding:"required"`
}
