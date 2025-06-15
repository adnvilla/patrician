package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

// GetDB initializes and returns the global gorm DB instance. It reads the
// connection string from the POSTGRES_DSN environment variable.
func GetDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		// fallback to default local connection for tests
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	}
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
