package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCommodityType(t *testing.T) {
	// Test commodity type constants
	assert.Equal(t, domain.CommodityType(1), domain.Load)
	assert.Equal(t, domain.CommodityType(2), domain.Barrel)

	// Test that the types are different
	assert.NotEqual(t, domain.Load, domain.Barrel)

	// Test that the types are ordered correctly
	assert.Less(t, int(domain.Load), int(domain.Barrel))

	// Test creating variables with these types
	var loadType domain.CommodityType = domain.Load
	var barrelType domain.CommodityType = domain.Barrel

	assert.Equal(t, domain.Load, loadType)
	assert.Equal(t, domain.Barrel, barrelType)
}

func TestCommodityTypeUsage(t *testing.T) {
	// Test that commodity types can be used in switch statements
	testType := domain.Barrel
	var result string

	switch testType {
	case domain.Load:
		result = "load"
	case domain.Barrel:
		result = "barrel"
	default:
		result = "unknown"
	}

	assert.Equal(t, "barrel", result)

	// Test with Load type
	testType = domain.Load
	switch testType {
	case domain.Load:
		result = "load"
	case domain.Barrel:
		result = "barrel"
	default:
		result = "unknown"
	}

	assert.Equal(t, "load", result)
}