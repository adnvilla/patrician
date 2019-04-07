package domain

// CommodityType Represents a type of commodities
type CommodityType int

// Commodities types
const (
	Load   CommodityType = iota + 1
	Barrel CommodityType = iota
)
