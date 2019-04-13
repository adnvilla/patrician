package domain

// Commodity Represents a commodity for trade
type Commodity struct {
	Entity
	Name          string
	CommodityType CommodityType
	Buy           int64
	Sell          int64
	Production    int64
	Consumption   int64
}

// GetStock Obtains Production result less Consumption
func (c *Commodity) GetStock() int64 {
	return c.Production - c.Consumption
}

// GetCommodities Obtains default commodities
func GetCommodities() map[string]*Commodity {
	return map[string]*Commodity{
		"Beer":      &Commodity{Name: "Beer", CommodityType: Barrel},
		"Bricks":    &Commodity{Name: "Bricks", CommodityType: Load},
		"Cloth":     &Commodity{Name: "Cloth", CommodityType: Barrel},
		"Fish":      &Commodity{Name: "Fish", CommodityType: Load},
		"Grain":     &Commodity{Name: "Grain", CommodityType: Load},
		"Hemp":      &Commodity{Name: "Hemp", CommodityType: Load},
		"Honey":     &Commodity{Name: "Honey", CommodityType: Barrel},
		"IronGoods": &Commodity{Name: "IronGoods", CommodityType: Barrel},
		"Leather":   &Commodity{Name: "Leather", CommodityType: Barrel},
		"Meat":      &Commodity{Name: "Meat", CommodityType: Load},
		"PigIron":   &Commodity{Name: "PigIron", CommodityType: Load},
		"Pitch":     &Commodity{Name: "Pitch", CommodityType: Barrel},
		"Pottery":   &Commodity{Name: "Pottery", CommodityType: Barrel},
		"Salt":      &Commodity{Name: "Salt", CommodityType: Barrel},
		"Skins":     &Commodity{Name: "Skins", CommodityType: Barrel},
		"Spices":    &Commodity{Name: "Spices", CommodityType: Barrel},
		"Timber":    &Commodity{Name: "Timber", CommodityType: Load},
		"WhaleOil":  &Commodity{Name: "WhaleOil", CommodityType: Barrel},
		"Wine":      &Commodity{Name: "Wine", CommodityType: Barrel},
		"Wool":      &Commodity{Name: "Wool", CommodityType: Load},
	}
}
