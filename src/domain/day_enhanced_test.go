package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestDayConstant(t *testing.T) {
	// Test Day constant value
	assert.Equal(t, float32(30), domain.Day)

	// Test that Day is a positive value
	assert.Greater(t, domain.Day, float32(0))

	// Test Day type
	assert.IsType(t, float32(0), domain.Day)
}

func TestWeekConstant(t *testing.T) {
	// Test Week constant value
	expectedWeek := domain.Day * 7
	assert.Equal(t, expectedWeek, domain.Week)
	assert.Equal(t, float32(210), domain.Week) // 30 * 7

	// Test that Week is a positive value
	assert.Greater(t, domain.Week, float32(0))

	// Test Week type
	assert.IsType(t, float32(0), domain.Week)
}

func TestDayWeekRelationship(t *testing.T) {
	// Test that Week equals Day * 7
	assert.Equal(t, domain.Day*7, domain.Week)

	// Test that Week is exactly 7 times Day
	assert.Equal(t, domain.Week/7, domain.Day)

	// Test mathematical relationship
	assert.InDelta(t, domain.Week, domain.Day*7, 0.0001)
}

func TestTimeConstants(t *testing.T) {
	// Test Day calculations
	halfDay := domain.Day / 2
	assert.Equal(t, float32(15), halfDay)

	doubleDay := domain.Day * 2
	assert.Equal(t, float32(60), doubleDay)

	// Test Week calculations
	halfWeek := domain.Week / 2
	assert.Equal(t, float32(105), halfWeek)

	twoWeeks := domain.Week * 2
	assert.Equal(t, float32(420), twoWeeks)
}

func TestTimeConstantsInCalculations(t *testing.T) {
	// Test Day in distance calculations (similar to what's used in City domain)
	distance := float32(3.0)
	distanceInDays := 2 * (distance / domain.Day)

	assert.InDelta(t, 0.2, distanceInDays, 0.0001) // 2 * (3 / 30) = 0.2

	// Test Week in supply calculations
	stock := int16(100)
	supplyCalculation := float32(stock) / (domain.Week / 30)
	assert.InDelta(t, 14.2857, supplyCalculation, 0.0001) // 100 / (210 / 30) = 100 / 7
}

func TestConstantImmutability(t *testing.T) {
	// Store original values
	originalDay := domain.Day
	originalWeek := domain.Week

	// Constants should remain unchanged after operations
	_ = domain.Day * 2
	_ = domain.Week / 3

	assert.Equal(t, originalDay, domain.Day)
	assert.Equal(t, originalWeek, domain.Week)
}

func TestTimeConstantPrecision(t *testing.T) {
	// Test float32 precision with Day
	result1 := domain.Day * 1.0
	result2 := float32(30.0)
	assert.Equal(t, result1, result2)

	// Test Week precision
	weekResult1 := domain.Week
	weekResult2 := float32(210.0)
	assert.Equal(t, weekResult1, weekResult2)
}

func TestTimeConstantsInContext(t *testing.T) {
	// Test how constants might be used in game calculations

	// Travel time calculation
	travelDistance := float32(2.5)
	travelTimeInDays := travelDistance / domain.Day
	assert.InDelta(t, 0.0833, travelTimeInDays, 0.0001)

	// Production cycle calculation
	productionPerDay := int16(10)
	productionPerWeek := float32(productionPerDay) * (domain.Week / domain.Day)
	assert.Equal(t, float32(70), productionPerWeek) // 10 * 7
}

func TestBoundaryValues(t *testing.T) {
	// Test very small multipliers
	verySmall := domain.Day * 0.001
	assert.Greater(t, verySmall, float32(0))
	assert.Equal(t, float32(0.03), verySmall)

	// Test very large multipliers
	veryLarge := domain.Day * 1000
	assert.Equal(t, float32(30000), veryLarge)

	// Test division edge cases
	halfOfSmallest := domain.Day / 2
	assert.Equal(t, float32(15), halfOfSmallest)
}

func TestConstantsConsistency(t *testing.T) {
	// Test that our constants make sense in the game context
	assert.Equal(t, float32(30), domain.Day)   // 30 seconds per game day
	assert.Equal(t, float32(210), domain.Week) // 210 seconds per game week

	// 7 days in a week should be consistent
	daysInWeek := domain.Week / domain.Day
	assert.InDelta(t, 7.0, daysInWeek, 0.0001)
}
