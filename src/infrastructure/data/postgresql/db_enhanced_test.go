package postgresql_test

import (
	"testing"

	"github.com/adnvilla/patrician/src/infrastructure/data/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestGetDBFunction(t *testing.T) {
	t.Run("GetDBWithoutPanickin", func(t *testing.T) {
		// Test that GetDB doesn't panic even if database is not available
		// This is important for testing environments

		// The function should not panic
		assert.NotPanics(t, func() {
			db, err := postgresql.GetDB()
			// We might get an error if no database is available, and that's okay
			_ = db
			_ = err
		})
	})

	t.Run("GetDBReturnsSomething", func(t *testing.T) {
		// Even if it fails, it should return something (even if it's an error)
		db, err := postgresql.GetDB()

		// Either we get a valid DB or an error, but not both nil
		if db == nil {
			assert.Error(t, err, "If DB is nil, there should be an error")
		}
		// If DB is not nil, error might still be present for connection issues
	})

	t.Run("ConsistentBehavior", func(t *testing.T) {
		// Multiple calls should behave consistently
		db1, err1 := postgresql.GetDB()
		db2, err2 := postgresql.GetDB()

		// If first call succeeded, second should return same instance
		if err1 == nil && db1 != nil {
			assert.Equal(t, db1, db2, "Should return same DB instance")
			assert.Equal(t, err1, err2, "Errors should be consistent")
		}
	})
}

func TestDatabaseConfiguration(t *testing.T) {
	t.Run("EnvironmentVariableHandling", func(t *testing.T) {
		// Test that the function handles environment variables appropriately
		// This is a behavioral test - we can't easily test the actual environment
		// variable reading without setting up complex mocking

		// The function should not panic regardless of environment setup
		assert.NotPanics(t, func() {
			postgresql.GetDB()
		})
	})
}

// Note: These tests are designed to work even when no actual database is available
// They test the code structure and error handling rather than actual database connectivity
// For integration tests with real database, see integration/ folder
