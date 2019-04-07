package domain

// City The city
type City struct {
	Entity
	Name       string
	MarketHall *MarketHall
}

func (c *City) getDistances() map[string]float64 {
	return Distances[c.Name]
}

func (c *City) UpdateCommodity(name string, buy, sell, production, consumption float64) {
	commodities := *c.MarketHall.Commodities
	commodity := commodities[name]

	commodity.Buy = buy
	commodity.Sell = sell
	commodity.Production = production
	commodity.Consumption = consumption
}

func (c *City) SetCommodities(commodities *map[string]Commodity) {
	c.MarketHall.Commodities = commodities
}

func (c *City) GetStockCommodity(name string) float64 {
	commodities := *c.MarketHall.Commodities
	commodity := commodities[name]

	return commodity.GetStock()
}

func (c *City) GetSupplyCommodityFromCity(namecommodity, namecity string) float64 {
	commodities := *c.MarketHall.Commodities
	commodity := commodities[namecommodity]

	distance := c.getDistances()[namecity]

	distanceInDays := 2 * (distance / Day)

	stock := commodity.GetStock()

	return (stock * distanceInDays) / Week
}

//Cities The Cities default
var Cities = map[string]City{
	"Edimburgo":   City{Name: "Edimburgo", MarketHall: &MarketHall{}},
	"Scarborough": City{Name: "Scarborough", MarketHall: &MarketHall{}},
	"Londres":     City{Name: "Londres", MarketHall: &MarketHall{}},
	"Brujas":      City{Name: "Brujas", MarketHall: &MarketHall{}},
	"Colonia":     City{Name: "Colonia", MarketHall: &MarketHall{}},
	"Groninga":    City{Name: "Groninga", MarketHall: &MarketHall{}},
	"Bremen":      City{Name: "Bremen", MarketHall: &MarketHall{}},
	"Hamburgo":    City{Name: "Hamburgo", MarketHall: &MarketHall{}},
	"Ripen":       City{Name: "Ripen", MarketHall: &MarketHall{}},
	"Bergen":      City{Name: "Bergen", MarketHall: &MarketHall{}},
	"Oslo":        City{Name: "Oslo", MarketHall: &MarketHall{}},
	"Aalborg":     City{Name: "Aalborg", MarketHall: &MarketHall{}},
	"Malmo":       City{Name: "Malmo", MarketHall: &MarketHall{}},
	"Lubeck":      City{Name: "Lubeck", MarketHall: &MarketHall{}},
	"Rostock":     City{Name: "Rostock", MarketHall: &MarketHall{}},
	"Stettin":     City{Name: "Stettin", MarketHall: &MarketHall{}},
	"Gdansk":      City{Name: "Gdansk", MarketHall: &MarketHall{}},
	"Torum":       City{Name: "Torum", MarketHall: &MarketHall{}},
	"Riga":        City{Name: "Riga", MarketHall: &MarketHall{}},
	"Visby":       City{Name: "Visby", MarketHall: &MarketHall{}},
	"Estocolmo":   City{Name: "Estocolmo", MarketHall: &MarketHall{}},
	"Reval":       City{Name: "Reval", MarketHall: &MarketHall{}},
	"Ladoga":      City{Name: "Ladoga", MarketHall: &MarketHall{}},
	"Novgorod":    City{Name: "Novgorod", MarketHall: &MarketHall{}},
}
