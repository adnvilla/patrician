package domain

// Commodity Represents a commodity for trade
type Commodity struct {
	Entity
	Name          string
	CommodityType CommodityType
	Buy           int16
	Sell          int16
	Production    int16
	Consumption   int16
}

// GetStock Obtains Production result less Consumption
func (c *Commodity) GetStock() int16 {
	return c.Production - c.Consumption
}

// GetCommodities Obtains default commodities
func GetCommodities() map[string]*Commodity {
	return map[string]*Commodity{
		"Beer":      {Name: "Beer", CommodityType: Barrel},
		"Bricks":    {Name: "Bricks", CommodityType: Load},
		"Cloth":     {Name: "Cloth", CommodityType: Barrel},
		"Fish":      {Name: "Fish", CommodityType: Load},
		"Grain":     {Name: "Grain", CommodityType: Load},
		"Hemp":      {Name: "Hemp", CommodityType: Load},
		"Honey":     {Name: "Honey", CommodityType: Barrel},
		"IronGoods": {Name: "IronGoods", CommodityType: Barrel},
		"Leather":   {Name: "Leather", CommodityType: Barrel},
		"Meat":      {Name: "Meat", CommodityType: Load},
		"PigIron":   {Name: "PigIron", CommodityType: Load},
		"Pitch":     {Name: "Pitch", CommodityType: Barrel},
		"Pottery":   {Name: "Pottery", CommodityType: Barrel},
		"Salt":      {Name: "Salt", CommodityType: Barrel},
		"Skins":     {Name: "Skins", CommodityType: Barrel},
		"Spices":    {Name: "Spices", CommodityType: Barrel},
		"Timber":    {Name: "Timber", CommodityType: Load},
		"WhaleOil":  {Name: "WhaleOil", CommodityType: Barrel},
		"Wine":      {Name: "Wine", CommodityType: Barrel},
		"Wool":      {Name: "Wool", CommodityType: Load},
	}
}
