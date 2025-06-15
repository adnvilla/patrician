package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func prepareCities() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/cities", handlers.GetCities)
	r.GET("/commodities", handlers.GetCommodities)
	r.GET("/distances", handlers.GetDistances)
	r.GET("/city/:name/commodities", handlers.GetCityCommodities)
	return r
}

func TestGetCitiesRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/cities", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetCommoditiesRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/commodities", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDistancesRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/distances", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCityCommoditiesRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/commodities", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
