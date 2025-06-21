package integration

import (
	"os"
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
)

func setupPostgresDB(t *testing.T) {
	db, err := postgresql.GetDB()
	if err != nil {
		t.Fatalf("connect db: %v", err)
	}
	err = db.AutoMigrate(&postgresql.CityModel{}, &postgresql.MarketHallModel{}, &postgresql.CommodityModel{}, &postgresql.DistanceModel{})
	if err != nil {
		t.Fatalf("migrate: %v", err)
	}
}

func TestCityRepositoryPostgres(t *testing.T) {
	if testing.Short() {
		t.Skip("short mode")
	}
	if os.Getenv("POSTGRES_DSN") == "" {
		t.Skip("POSTGRES_DSN not set")
	}
	setupPostgresDB(t)
	db, _ := postgresql.GetDB()
	repo := postgresql.NewCityRepository(db)

	commodities := domain.GetCommodities()
	mh := domain.MarketHall{Commodities: commodities}
	city := domain.City{Name: "IntegrationCity", MarketHall: mh}

	if err := repo.Create(&city); err != nil {
		t.Fatalf("create: %v", err)
	}

	cities, err := repo.FindAll()
	if err != nil || len(cities) == 0 {
		t.Fatalf("findall: %v", err)
	}
}
