package domain

// Commodity Represents a commodity for trade
type Commodity struct {
	Entity
	CommodityType
	Buy         int64
	Sell        int64
	Production  int64
	Consumption int64
}

// GetStock Obtains Production result less Consumption
func (c *Commodity) GetStock() int64 {
	return c.Production - c.Consumption
}

// GetCommodities Obtains default commodities
func GetCommodities() map[string]*Commodity {
	return map[string]*Commodity{
		"Beer":      &Commodity{CommodityType: Barrel},
		"Bricks":    &Commodity{CommodityType: Load},
		"Cloth":     &Commodity{CommodityType: Barrel},
		"Fish":      &Commodity{CommodityType: Load},
		"Grain":     &Commodity{CommodityType: Load},
		"Hemp":      &Commodity{CommodityType: Load},
		"Honey":     &Commodity{CommodityType: Barrel},
		"IronGoods": &Commodity{CommodityType: Barrel},
		"Leather":   &Commodity{CommodityType: Barrel},
		"Meat":      &Commodity{CommodityType: Load},
		"PigIron":   &Commodity{CommodityType: Load},
		"Pitch":     &Commodity{CommodityType: Barrel},
		"Pottery":   &Commodity{CommodityType: Barrel},
		"Salt":      &Commodity{CommodityType: Barrel},
		"Skins":     &Commodity{CommodityType: Barrel},
		"Spices":    &Commodity{CommodityType: Barrel},
		"Timber":    &Commodity{CommodityType: Load},
		"WhaleOil":  &Commodity{CommodityType: Barrel},
		"Wine":      &Commodity{CommodityType: Barrel},
		"Wool":      &Commodity{CommodityType: Load},
	}
}
