package domain

// City The city
type City struct {
	Entity
	Name       string
	MarketHall MarketHall
}

func (c *City) getDistances() map[string]float64 {
	return Distances[c.Name]
}

func (c *City) UpdateCommodity(name string, buy, sell, production, consumption int64) {
	commodities := c.MarketHall.Commodities
	commodity := commodities[name]

	commodity.Buy = buy
	commodity.Sell = sell
	commodity.Production = production
	commodity.Consumption = consumption
}

func (c *City) GetCommodities() map[string]*Commodity {
	return c.MarketHall.Commodities
}

func (c *City) SetCommodities(commodities map[string]*Commodity) {
	c.MarketHall.Commodities = commodities
}

func (c *City) SetMarketHall(markethall MarketHall) {
	c.MarketHall = markethall
}

func (c *City) GetStockCommodity(name string) int64 {
	commodities := c.MarketHall.Commodities
	commodity := commodities[name]

	return commodity.GetStock()
}

func (c *City) GetSupplyCommodityFromCity(namecommodity, namecity string) int64 {
	commodities := c.MarketHall.Commodities
	commodity := commodities[namecommodity]

	distance := c.getDistances()[namecity]

	distanceInDays := 2 * (distance / Day)

	stock := commodity.GetStock()

	return int64((float64(stock) * distanceInDays) / (Week / 30))
}

func (c *City) GetSupplyCommoditiesFromCity(city string) map[string]int64 {
	supply := map[string]int64{}
	for name := range c.MarketHall.Commodities {
		supply[name] = c.GetSupplyCommodityFromCity(name, city)
	}
	return supply
}

func (c *City) GetStockCommodities() map[string]int64 {
	stocks := map[string]int64{}
	for name := range c.MarketHall.Commodities {
		stocks[name] = c.GetStockCommodity(name)
	}
	return stocks
}

//Cities The Cities default
var Cities = map[string]*City{
	"Edimburgo":   {Name: "Edimburgo"},
	"Scarborough": {Name: "Scarborough"},
	"Londres":     {Name: "Londres"},
	"Brujas":      {Name: "Brujas"},
	"Colonia":     {Name: "Colonia"},
	"Groninga":    {Name: "Groninga"},
	"Bremen":      {Name: "Bremen"},
	"Hamburgo":    {Name: "Hamburgo"},
	"Ripen":       {Name: "Ripen"},
	"Bergen":      {Name: "Bergen"},
	"Oslo":        {Name: "Oslo"},
	"Aalborg":     {Name: "Aalborg"},
	"Malmo":       {Name: "Malmo"},
	"Lubeck":      {Name: "Lubeck"},
	"Rostock":     {Name: "Rostock"},
	"Stettin":     {Name: "Stettin"},
	"Gdansk":      {Name: "Gdansk"},
	"Torum":       {Name: "Torum"},
	"Riga":        {Name: "Riga"},
	"Visby":       {Name: "Visby"},
	"Estocolmo":   {Name: "Estocolmo"},
	"Reval":       {Name: "Reval"},
	"Ladoga":      {Name: "Ladoga"},
	"Novgorod":    {Name: "Novgorod"},
}
