package controllers

import (
	"net/http"

	usecases "github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	usecase := usecases.NewGetCitiesUseCase()

	result, _ := usecase.Handle(c, nil)

	c.JSON(http.StatusOK, result)
}

func GetCityCommodities(c *gin.Context) {
	usecase := usecases.NewGetCityCommoditiesUseCase()

	name := c.Param("name")
	input := usecases.GetCityCommoditiesInput{
		CityName: name,
	}

	result, _ := usecase.Handle(c, input)

	c.JSON(http.StatusOK, result)
}
