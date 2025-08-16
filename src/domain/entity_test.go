package domain_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	// Create an Entity instance
	entity := domain.Entity{}

	// Test that the ID is zero value initially
	assert.Equal(t, uint(0), entity.ID)

	// Test that CreatedAt, UpdatedAt, and DeletedAt fields exist from gorm.Model
	// These are implicitly tested by ensuring the struct can be created and used
	assert.NotPanics(t, func() {
		_ = entity.CreatedAt
		_ = entity.UpdatedAt
		_ = entity.DeletedAt
	})
}
