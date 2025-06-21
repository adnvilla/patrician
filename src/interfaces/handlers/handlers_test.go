package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/getkin/kin-openapi/openapi3"
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
	swagger, err := openapi3.NewLoader().LoadFromFile("../../../docs/apenapi.yaml")
	if err != nil {
		panic(err)
	}
	if err = swagger.Validate(context.Background()); err != nil {
		panic(err)
	}
	r.Use(middleware.OapiRequestValidator(swagger))
	r.GET("/cities", handlers.GetCities)
	r.GET("/commodities", handlers.GetCommodities)
	r.GET("/distances", handlers.GetDistances)
	r.GET("/city/:name/commodities", handlers.GetCityCommodities)
	r.POST("/city/:name/commodity", handlers.UpdateCommodity)
	r.POST("/city/:name/commodities", handlers.UpdateCommodities)
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

func TestUpdateCommodityValidation(t *testing.T) {
	prepareCities()
	router := setupRouter()

	body := `{"name":"Beer","buy":1}`
	req, _ := http.NewRequest(http.MethodPost, "/city/Edimburgo/commodity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateCommodityOK(t *testing.T) {
	prepareCities()
	router := setupRouter()

	body := `{"name":"Beer","buy":1,"sell":1,"production":1,"consumption":1}`
	req, _ := http.NewRequest(http.MethodPost, "/city/Edimburgo/commodity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
