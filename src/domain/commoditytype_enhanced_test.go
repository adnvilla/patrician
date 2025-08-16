package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCommodityTypeConstants(t *testing.T) {
	// Test that Load is 1 and Barrel is 2
	assert.Equal(t, domain.CommodityType(1), domain.Load)
	assert.Equal(t, domain.CommodityType(2), domain.Barrel)

	// Test that they are different values
	assert.NotEqual(t, domain.Load, domain.Barrel)
}

func TestCommodityTypeValues(t *testing.T) {
	// Test Load value
	assert.Equal(t, int(1), int(domain.Load))

	// Test Barrel value
	assert.Equal(t, int(2), int(domain.Barrel))
}

func TestCommodityTypeOrdering(t *testing.T) {
	// Test that Load comes before Barrel in the iota sequence
	assert.Less(t, int(domain.Load), int(domain.Barrel))

	// Test specific ordering
	assert.Equal(t, int(domain.Load)+1, int(domain.Barrel))
}

func TestCommodityTypeComparisons(t *testing.T) {
	// Test equality
	assert.Equal(t, domain.Load, domain.Load)
	assert.Equal(t, domain.Barrel, domain.Barrel)

	// Test inequality
	assert.NotEqual(t, domain.Load, domain.Barrel)

	// Test comparison operators
	assert.True(t, domain.Load < domain.Barrel)
	assert.False(t, domain.Load > domain.Barrel)
	assert.True(t, domain.Barrel > domain.Load)
	assert.False(t, domain.Barrel < domain.Load)
}

func TestCommodityTypeInCommodities(t *testing.T) {
	commodities := domain.GetCommodities()

	loadCommodities := []string{}
	barrelCommodities := []string{}

	for name, commodity := range commodities {
		switch commodity.CommodityType {
		case domain.Load:
			loadCommodities = append(loadCommodities, name)
		case domain.Barrel:
			barrelCommodities = append(barrelCommodities, name)
		}
	}

	// Test that we have both types
	assert.Greater(t, len(loadCommodities), 0, "Should have Load type commodities")
	assert.Greater(t, len(barrelCommodities), 0, "Should have Barrel type commodities")

	// Test specific Load commodities
	expectedLoadCommodities := []string{"Bricks", "Fish", "Grain", "Hemp", "Meat", "PigIron", "Timber", "Wool"}
	for _, name := range expectedLoadCommodities {
		commodity := commodities[name]
		assert.Equal(t, domain.Load, commodity.CommodityType, "Commodity %s should be Load type", name)
	}

	// Test specific Barrel commodities
	expectedBarrelCommodities := []string{"Beer", "Cloth", "Honey", "IronGoods", "Leather", "Pitch", "Pottery", "Salt", "Skins", "Spices", "WhaleOil", "Wine"}
	for _, name := range expectedBarrelCommodities {
		commodity := commodities[name]
		assert.Equal(t, domain.Barrel, commodity.CommodityType, "Commodity %s should be Barrel type", name)
	}
}

func TestCommodityTypeAsInt(t *testing.T) {
	// Test conversion to int
	loadInt := int(domain.Load)
	barrelInt := int(domain.Barrel)

	assert.Equal(t, 1, loadInt)
	assert.Equal(t, 2, barrelInt)

	// Test that conversion is consistent
	assert.Equal(t, loadInt+1, barrelInt)
}

func TestCommodityTypeFromInt(t *testing.T) {
	// Test conversion from int
	loadFromInt := domain.CommodityType(1)
	barrelFromInt := domain.CommodityType(2)

	assert.Equal(t, domain.Load, loadFromInt)
	assert.Equal(t, domain.Barrel, barrelFromInt)
}

func TestCommodityTypeZeroValue(t *testing.T) {
	var zeroType domain.CommodityType

	// Zero value should be 0 (not Load or Barrel)
	assert.Equal(t, domain.CommodityType(0), zeroType)
	assert.NotEqual(t, domain.Load, zeroType)
	assert.NotEqual(t, domain.Barrel, zeroType)
}

func TestCommodityTypeInStructs(t *testing.T) {
	// Test Load commodity
	loadCommodity := domain.Commodity{
		Name:          "TestLoad",
		CommodityType: domain.Load,
	}
	assert.Equal(t, domain.Load, loadCommodity.CommodityType)

	// Test Barrel commodity
	barrelCommodity := domain.Commodity{
		Name:          "TestBarrel",
		CommodityType: domain.Barrel,
	}
	assert.Equal(t, domain.Barrel, barrelCommodity.CommodityType)
}

func TestCommodityTypeSwitch(t *testing.T) {
	testCases := []struct {
		commodityType domain.CommodityType
		expected      string
	}{
		{domain.Load, "Load"},
		{domain.Barrel, "Barrel"},
	}

	for _, tc := range testCases {
		var result string
		switch tc.commodityType {
		case domain.Load:
			result = "Load"
		case domain.Barrel:
			result = "Barrel"
		default:
			result = "Unknown"
		}
		assert.Equal(t, tc.expected, result)
	}
}

func TestCommodityTypeDistribution(t *testing.T) {
	commodities := domain.GetCommodities()

	typeCount := make(map[domain.CommodityType]int)
	for _, commodity := range commodities {
		typeCount[commodity.CommodityType]++
	}

	// Should have both types represented
	assert.Greater(t, typeCount[domain.Load], 0, "Should have Load commodities")
	assert.Greater(t, typeCount[domain.Barrel], 0, "Should have Barrel commodities")

	// Total should equal all commodities
	total := typeCount[domain.Load] + typeCount[domain.Barrel]
	assert.Equal(t, len(commodities), total)
}
