package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestDayWeekConstants(t *testing.T) {
	assert.Equal(t, float32(30), domain.Day)
	assert.Equal(t, float32(210), domain.Week)
}
