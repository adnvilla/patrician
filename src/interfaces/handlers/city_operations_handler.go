package handlers

import (
	"net/http"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/gin-gonic/gin"
)

func GetStock(c *gin.Context) {
	name := c.Param("name")

	city := domain.Cities[name]

	supply := city.GetStockCommodities()

	c.JSON(http.StatusOK, supply)
}

func GetSupply(c *gin.Context) {
	name := c.Param("name")
	from := c.Param("city")

	city := domain.Cities[name]

	supply := city.GetSupplyCommoditiesFromCity(from)

	c.JSON(http.StatusOK, supply)
}

func UpdateCommodity(c *gin.Context) {
	name := c.Param("name")

	var commodity Commodity

	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := domain.Cities[name]

	if err := city.UpdateCommodity(commodity.Name, commodity.Buy, commodity.Sell, commodity.Production, commodity.Consumption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "updated!")
}

func UpdateCommodities(c *gin.Context) {
	name := c.Param("name")

	var commodities Commodities

	if err := c.ShouldBindJSON(&commodities); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := domain.Cities[name]

	for _, commodity := range commodities.Commodities {
		if err := city.UpdateCommodity(commodity.Name, commodity.Buy, commodity.Sell, commodity.Production, commodity.Consumption); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.String(http.StatusOK, "updated!")
}

// Commodity represents a tradeable good that can be updated.
type Commodity struct {
	Name        string `json:"name" binding:"required"`
	Buy         int16  `json:"buy" binding:"required"`
	Sell        int16  `json:"sell" binding:"required"`
	Production  int16  `json:"production" binding:"required"`
	Consumption int16  `json:"consumption" binding:"required"`
}

// Commodities is a helper wrapper for bulk updates of commodities.
type Commodities struct {
	Commodities []Commodity `json:"commodities" binding:"required"`
}
