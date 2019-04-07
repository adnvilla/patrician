package domain

// City The city
type City struct {
	Entity
	Name       string
	MarketHall MarketHall
}

func (c *City) getDistances() map[string]int64 {
	return Distances[c.Name]
}

func (c *City) UpdateCommodity(name string, buy, sell, production, consumption float64) {
	commodity := c.MarketHall.Commodities[name]

	commodity.Buy = buy
	commodity.Sell = sell
	commodity.Production = production
	commodity.Consumption = consumption
}

func (c *City) SetCommodities(commodities map[string]Commodity) {
	c.MarketHall.Commodities = commodities
}
