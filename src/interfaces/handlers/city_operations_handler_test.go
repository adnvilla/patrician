package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGetStockRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	// Add the GetStock route
	router.GET("/city/:name/stock", handlers.GetStock)

	req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/stock", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestGetSupplyRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	// Add the GetSupply route
	router.GET("/city/:name/supply/:city", handlers.GetSupply)

	req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/supply/Visby", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}

func TestUpdateCommodityRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	// Create the payload
	commodity := handlers.Commodity{
		Name:        "Beer",
		Buy:         15,
		Sell:        20,
		Production:  10,
		Consumption: 100,
	}
	payload, _ := json.Marshal(commodity)

	req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodity", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "updated!", w.Body.String())

	// Verify the update was actually made
	city := domain.Cities["Estocolmo"]
	updatedCommodity := city.MarketHall.Commodities["Beer"]
	assert.Equal(t, int16(15), updatedCommodity.Buy)
	assert.Equal(t, int16(20), updatedCommodity.Sell)
}

func TestUpdateCommoditiesRoute(t *testing.T) {
	prepareCities()
	router := setupRouter()

	// Create the payload
	commodities := handlers.Commodities{
		Commodities: []handlers.Commodity{
			{
				Name:        "Beer",
				Buy:         16,
				Sell:        21,
				Production:  11,
				Consumption: 101,
			},
			{
				Name:        "Fish",
				Buy:         10,
				Sell:        15,
				Production:  8,
				Consumption: 80,
			},
		},
	}
	payload, _ := json.Marshal(commodities)

	req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodities", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "updated!", w.Body.String())

	// Verify the updates were actually made
	city := domain.Cities["Estocolmo"]
	updatedBeer := city.MarketHall.Commodities["Beer"]
	assert.Equal(t, int16(16), updatedBeer.Buy)
	assert.Equal(t, int16(21), updatedBeer.Sell)

	updatedFish := city.MarketHall.Commodities["Fish"]
	assert.Equal(t, int16(10), updatedFish.Buy)
	assert.Equal(t, int16(15), updatedFish.Sell)
}

func TestUpdateCommodityBadRequest(t *testing.T) {
	prepareCities()
	router := setupRouter()

	// Create an invalid payload
	payload := []byte(`{"name": "Beer", "buy": "invalid"}`) // buy should be an int

	req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodity", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateCommodityNotFound(t *testing.T) {
	prepareCities()
	router := setupRouter()

	// Create a payload with non-existent commodity
	commodity := handlers.Commodity{
		Name:        "NonExistentCommodity",
		Buy:         15,
		Sell:        20,
		Production:  10,
		Consumption: 100,
	}
	payload, _ := json.Marshal(commodity)

	req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodity", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
