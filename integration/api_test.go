package integration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/adnvilla/patrician/src/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

// setupRouter configures gin router with handlers.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/cities", handlers.GetCities)
	r.GET("/commodities", handlers.GetCommodities)
	r.GET("/distances", handlers.GetDistances)
	r.GET("/city/:name/commodities", handlers.GetCityCommodities)
	return r
}

// prepareCities resets domain data before each test.
func prepareCities() {
	for _, city := range domain.Cities {
		commodities := domain.GetCommodities()
		city.SetMarketHall(domain.MarketHall{Commodities: commodities})
	}
}

func setupDB(t *testing.T) {
	db, err := postgresql.GetDB()
	if err != nil {
		t.Fatalf("failed to connect db: %v", err)
	}
	err = db.AutoMigrate(&postgresql.CityModel{}, &postgresql.MarketHallModel{}, &postgresql.CommodityModel{}, &postgresql.DistanceModel{})
	if err != nil {
		t.Fatalf("migrate: %v", err)
	}
}

func TestCitiesEndpoint(t *testing.T) {
	if os.Getenv("POSTGRES_DSN") == "" {
		t.Skip("POSTGRES_DSN not set")
	}
	setupDB(t)
	prepareCities()
	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/cities", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", w.Code)
	}
}
