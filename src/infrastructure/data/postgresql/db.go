package postgresql

import (
        "os"

        dbgo "github.com/adnvilla/db-go"
        "gorm.io/gorm"
)

// GetDB initializes and returns the global gorm DB instance. It reads the
// connection string from the POSTGRES_DSN environment variable and uses the
// adnvilla/db-go package to establish the connection.
func GetDB() (*gorm.DB, error) {
        dsn := os.Getenv("POSTGRES_DSN")
        if dsn == "" {
                // fallback to default local connection for tests
                dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
        }

        conn := dbgo.GetConnection(dbgo.Config{PrimaryDSN: dsn})
        return conn.Instance, conn.Error
}
