package domain

type MarketHall struct {
	Entity
	Commodities map[string]*Commodity `gorm:"-"`
}
