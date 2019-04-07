package main

import (
	"github.com/adnvilla/patrician/patrician/domain"
)

func main() {
	for _, city := range Cities {
		city.SetCommodities(GetCommodities())
	}
}

//Cities The Cities default
var Cities = map[string]domain.City{
	"Edimburgo":   domain.City{Name: "Edimburgo", MarketHall: domain.MarketHall{}},
	"Scarborough": domain.City{Name: "Scarborough", MarketHall: domain.MarketHall{}},
	"Londres":     domain.City{Name: "Londres", MarketHall: domain.MarketHall{}},
	"Brujas":      domain.City{Name: "Brujas", MarketHall: domain.MarketHall{}},
	"Colonia":     domain.City{Name: "Colonia", MarketHall: domain.MarketHall{}},
	"Groninga":    domain.City{Name: "Groninga", MarketHall: domain.MarketHall{}},
	"Bremen":      domain.City{Name: "Bremen", MarketHall: domain.MarketHall{}},
	"Hamburgo":    domain.City{Name: "Hamburgo", MarketHall: domain.MarketHall{}},
	"Ripen":       domain.City{Name: "Ripen", MarketHall: domain.MarketHall{}},
	"Bergen":      domain.City{Name: "Bergen", MarketHall: domain.MarketHall{}},
	"Oslo":        domain.City{Name: "Oslo", MarketHall: domain.MarketHall{}},
	"Aalborg":     domain.City{Name: "Aalborg", MarketHall: domain.MarketHall{}},
	"Malmo":       domain.City{Name: "Malmo", MarketHall: domain.MarketHall{}},
	"Lubeck":      domain.City{Name: "Lubeck", MarketHall: domain.MarketHall{}},
	"Rostock":     domain.City{Name: "Rostock", MarketHall: domain.MarketHall{}},
	"Stettin":     domain.City{Name: "Stettin", MarketHall: domain.MarketHall{}},
	"Gdansk":      domain.City{Name: "Gdansk", MarketHall: domain.MarketHall{}},
	"Torum":       domain.City{Name: "Torum", MarketHall: domain.MarketHall{}},
	"Riga":        domain.City{Name: "Riga", MarketHall: domain.MarketHall{}},
	"Visby":       domain.City{Name: "Visby", MarketHall: domain.MarketHall{}},
	"Estocolmo":   domain.City{Name: "Estocolmo", MarketHall: domain.MarketHall{}},
	"Reval":       domain.City{Name: "Reval", MarketHall: domain.MarketHall{}},
	"Ladoga":      domain.City{Name: "Ladoga", MarketHall: domain.MarketHall{}},
	"Novgorod":    domain.City{Name: "Novgorod", MarketHall: domain.MarketHall{}},
}

// GetCommodities Obtains default commodities
func GetCommodities() map[string]domain.Commodity {
	return map[string]domain.Commodity{
		"Beer":      domain.Commodity{CommodityType: domain.Barrel},
		"Bricks":    domain.Commodity{CommodityType: domain.Load},
		"Cloth":     domain.Commodity{CommodityType: domain.Barrel},
		"Fish":      domain.Commodity{CommodityType: domain.Load},
		"Grain":     domain.Commodity{CommodityType: domain.Load},
		"Hemp":      domain.Commodity{CommodityType: domain.Load},
		"Honey":     domain.Commodity{CommodityType: domain.Barrel},
		"IronGoods": domain.Commodity{CommodityType: domain.Barrel},
		"Leather":   domain.Commodity{CommodityType: domain.Barrel},
		"Meat":      domain.Commodity{CommodityType: domain.Load},
		"PigIron":   domain.Commodity{CommodityType: domain.Load},
		"Pitch":     domain.Commodity{CommodityType: domain.Barrel},
		"Pottery":   domain.Commodity{CommodityType: domain.Barrel},
		"Salt":      domain.Commodity{CommodityType: domain.Barrel},
		"Skins":     domain.Commodity{CommodityType: domain.Barrel},
		"Spices":    domain.Commodity{CommodityType: domain.Barrel},
		"Timber":    domain.Commodity{CommodityType: domain.Load},
		"WhaleOil":  domain.Commodity{CommodityType: domain.Barrel},
		"Wine":      domain.Commodity{CommodityType: domain.Barrel},
		"Wool":      domain.Commodity{CommodityType: domain.Load},
	}
}
