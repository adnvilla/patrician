package handlers

import (
	"net/http"

	usecases "github.com/adnvilla/patrician/internal/application/use_cases"
	"github.com/gin-gonic/gin"
)

func GetCommodities(c *gin.Context) {
	usecase := usecases.NewGetCommoditiesUseCase()

	result, _ := usecase.Handle(c, nil)

	c.JSON(http.StatusOK, result)
}
