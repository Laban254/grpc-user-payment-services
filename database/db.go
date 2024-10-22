package database

import (
    "log"

    "gorm.io/driver/postgres" // Import the GORM PostgreSQL driver
    "gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection using GORM
func ConnectDB() {
    var err error
	connStr := "host=localhost port=5432 user=grpc password=password dbname=grpc sslmode=disable"

    
    // Use gorm.Open with the PostgreSQL driver
    DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Optional: Ping the database (not needed for GORM, but you can run a simple query if you want)
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get DB instance: %v", err)
    }
    
    if err = sqlDB.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
	// Auto migrate the User model
    if err := DB.AutoMigrate(&User{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    
    log.Println("Connected to PostgreSQL database")
}
