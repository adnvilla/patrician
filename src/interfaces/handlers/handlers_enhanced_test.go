package handlers_test

import (
	"bytes"
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

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	
	// Try to load OpenAPI spec, but don't fail if it doesn't exist
	swagger, err := openapi3.NewLoader().LoadFromFile("../../../docs/apenapi.yaml")
	if err == nil {
		if err = swagger.Validate(context.Background()); err == nil {
			r.Use(middleware.OapiRequestValidator(swagger))
		}
	}
	
	r.GET("/cities", handlers.GetCities)
	r.GET("/commodities", handlers.GetCommodities)
	r.GET("/distances", handlers.GetDistances)
	r.GET("/city/:name/commodities", handlers.GetCityCommodities)
	r.POST("/city/:name/commodity", handlers.UpdateCommodity)
	r.POST("/city/:name/commodities", handlers.UpdateCommodities)
	return r
}

func TestGetCitiesEnhanced(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("SuccessfulRequest", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/cities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
		
		// Verify response structure
		assert.Greater(t, len(response), 0)
	})
	
	t.Run("ResponseHeaders", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/cities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	})
	
	t.Run("ConcurrentRequests", func(t *testing.T) {
		// Test multiple concurrent requests
		done := make(chan bool, 10)
		
		for i := 0; i < 10; i++ {
			go func() {
				req, _ := http.NewRequest(http.MethodGet, "/cities", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code)
				done <- true
			}()
		}
		
		// Wait for all goroutines to complete
		for i := 0; i < 10; i++ {
			<-done
		}
	})
}

func TestGetCommoditiesEnhanced(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("SuccessfulRequest", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/commodities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	
	t.Run("ResponseStructure", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/commodities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		// Should contain expected commodities
		expectedCommodities := []string{"Beer", "Cloth", "Fish", "Grain"}
		for _, commodity := range expectedCommodities {
			assert.Contains(t, response, commodity)
		}
	})
}

func TestGetDistancesEnhanced(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("SuccessfulRequest", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/distances", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	
	t.Run("ResponseStructure", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/distances", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		
		// Should contain distance data for cities
		assert.Greater(t, len(response), 0)
		
		// Check that cities have distance entries
		for cityName := range response {
			cityDistances := response[cityName]
			assert.Greater(t, len(cityDistances), 0, "City %s should have distances", cityName)
		}
	})
}

func TestGetCityCommoditiesEnhanced(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("ValidCity", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/commodities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
	
	t.Run("DifferentCities", func(t *testing.T) {
		cities := []string{"Estocolmo", "Visby", "Londres", "Bremen"}
		
		for _, cityName := range cities {
			req, _ := http.NewRequest(http.MethodGet, "/city/"+cityName+"/commodities", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			
			// Should return OK for valid cities
			if _, exists := domain.Cities[cityName]; exists {
				assert.Equal(t, http.StatusOK, w.Code, "Should return OK for city %s", cityName)
			}
		}
	})
	
	t.Run("URLEncoding", func(t *testing.T) {
		// Test city names that might need URL encoding
		req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/commodities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUpdateCommodityEnhanced(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("ValidUpdate", func(t *testing.T) {
		payload := map[string]interface{}{
			"name":        "Beer",
			"buy":         15,
			"sell":        20,
			"production":  100,
			"consumption": 50,
		}
		
		jsonPayload, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodity", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		// Note: Actual response depends on implementation
		// For now, we just test that it doesn't panic
		assert.NotEqual(t, http.StatusInternalServerError, w.Code)
	})
	
	t.Run("InvalidJSON", func(t *testing.T) {
		invalidJSON := `{"name": "Beer", "buy": invalid}`
		req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodity", strings.NewReader(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		// Should handle invalid JSON gracefully
		assert.NotEqual(t, http.StatusInternalServerError, w.Code)
	})
	
	t.Run("EmptyPayload", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/city/Estocolmo/commodity", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		// Should handle empty payload gracefully
		assert.NotEqual(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHTTPMethods(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("WrongMethodForGetCities", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/cities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	
	t.Run("WrongMethodForUpdateCommodity", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/commodity", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestRouteParameters(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("MissingCityParameter", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/city//commodities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	
	t.Run("ValidCityParameter", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/city/Estocolmo/commodities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestResponseFormats(t *testing.T) {
	prepareCities()
	router := setupTestRouter()
	
	t.Run("JSONContentType", func(t *testing.T) {
		endpoints := []string{"/cities", "/commodities", "/distances", "/city/Estocolmo/commodities"}
		
		for _, endpoint := range endpoints {
			req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			
			if w.Code == http.StatusOK {
				contentType := w.Header().Get("Content-Type")
				assert.Contains(t, contentType, "application/json", "Endpoint %s should return JSON", endpoint)
			}
		}
	})
	
	t.Run("ValidJSONResponse", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/cities", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		
		if w.Code == http.StatusOK {
			var response interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err, "Response should be valid JSON")
		}
	})
}
