package postgresql

import (
	"github.com/adnvilla/patrician/src/domain"
)

// CityModel mirrors domain.City for persistence.
type CityModel struct {
	domain.Entity
	Name         string
	MarketHallID uint
	MarketHall   MarketHallModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:MarketHallID"`
}

type CommodityModel struct {
	domain.Entity
	Name          string
	CommodityType int
	Buy           int16
	Sell          int16
	Production    int16
	Consumption   int16
	MarketHallID  uint
}

// MarketHallModel stores the commodities of a city.
type MarketHallModel struct {
	domain.Entity
	Commodities []CommodityModel `gorm:"foreignKey:MarketHallID;constraint:OnDelete:CASCADE"`
	CityID      uint
}

// DistanceModel represents the distance between two cities.
type DistanceModel struct {
	domain.Entity
	FromCity string
	ToCity   string
	Value    float32
}
