package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	err = db.AutoMigrate(&postgresql.CityModel{}, &postgresql.MarketHallModel{}, &postgresql.CommodityModel{}, &postgresql.DistanceModel{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestCityRepositoryCRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCityRepository(db)

	commodities := domain.GetCommodities()
	mh := domain.MarketHall{Commodities: commodities}
	city := domain.City{Name: "TestCity", MarketHall: mh}

	if err := repo.Create(&city); err != nil {
		t.Fatalf("create: %v", err)
	}

	cities, err := repo.FindAll()
	if err != nil || len(cities) != 1 {
		t.Fatalf("findall error: %v len=%d", err, len(cities))
	}

	city.Name = "UpdatedCity"
	if err := repo.Update(&city); err != nil {
		t.Fatalf("update: %v", err)
	}

	fetched, err := repo.FindByID(cities[0].ID)
	if err != nil || fetched.Name != "UpdatedCity" {
		t.Fatalf("findbyid: %v name=%s", err, fetched.Name)
	}

	if err := repo.Delete(&city); err != nil {
		t.Fatalf("delete: %v", err)
	}

	var count int64
	db.Model(&postgresql.CityModel{}).Count(&count)
	if count != 0 {
		t.Fatalf("expected soft delete count 0 got %d", count)
	}
}

func TestCommodityRepositoryCRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCommodityRepository(db)

	c := domain.Commodity{Name: "Beer", Buy: 1, Sell: 2, Production: 3, Consumption: 1}
	if err := repo.Create(&c); err != nil {
		t.Fatalf("create: %v", err)
	}
	c.Sell = 5
	if err := repo.Update(&c); err != nil {
		t.Fatalf("update: %v", err)
	}
	fetched, err := repo.FindByID(c.ID)
	if err != nil || fetched.Sell != 5 {
		t.Fatalf("find: %v sell=%d", err, fetched.Sell)
	}
	if err := repo.Delete(&c); err != nil {
		t.Fatalf("delete: %v", err)
	}
	var count int64
	db.Model(&postgresql.CommodityModel{}).Count(&count)
	if count != 0 {
		t.Fatalf("expected 0 got %d", count)
	}
}

func TestDistanceRepositoryCRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewDistanceRepository(db)

	d := domain.Distance{FromCity: "A", ToCity: "B", Value: 2}
	if err := repo.Create(&d); err != nil {
		t.Fatalf("create: %v", err)
	}
	d.Value = 3
	if err := repo.Update(&d); err != nil {
		t.Fatalf("update: %v", err)
	}
	list, err := repo.FindAll()
	if err != nil || len(list) != 1 || list[0].Value != 3 {
		t.Fatalf("findall: %v", err)
	}
	if err := repo.Delete(&d); err != nil {
		t.Fatalf("delete: %v", err)
	}
	var count int64
	db.Model(&postgresql.DistanceModel{}).Count(&count)
	if count != 0 {
		t.Fatalf("expected 0 got %d", count)
	}
}

func TestMarketHallRepositoryCRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewMarketHallRepository(db)

	commodities := domain.GetCommodities()
	mh := domain.MarketHall{Commodities: commodities}

	if err := repo.Create(&mh); err != nil {
		t.Fatalf("create: %v", err)
	}
	mh.Commodities["Beer"].Sell = 11
	if err := repo.Update(&mh); err != nil {
		t.Fatalf("update: %v", err)
	}
	fetched, err := repo.FindByID(mh.ID)
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if fetched.Commodities["Beer"].Sell != 11 {
		t.Fatalf("expected sell 11 got %d", fetched.Commodities["Beer"].Sell)
	}
	if err := repo.Delete(&mh); err != nil {
		t.Fatalf("delete: %v", err)
	}
	var count int64
	db.Model(&postgresql.MarketHallModel{}).Count(&count)
	if count != 0 {
		t.Fatalf("expected 0 got %d", count)
	}
}
