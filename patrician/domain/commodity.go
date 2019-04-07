package domain

// Commodity Represents a commodity for trade
type Commodity struct {
	Entity
	CommodityType
	Buy         float64
	Sell        float64
	Production  float64
	Consumption float64
}
