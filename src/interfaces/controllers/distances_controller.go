package controllers

import (
	"net/http"

	usecases "github.com/adnvilla/patrician/src/application/use_cases"
	"github.com/gin-gonic/gin"
)

func GetDistances(c *gin.Context) {
	usecase := usecases.NewGetDistancesUseCase()

	result, _ := usecase.Handle(c, nil)

	c.JSON(http.StatusOK, result)
}
