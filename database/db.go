package database

import (
    "log"

    "gorm.io/driver/postgres" 
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
	connStr := "host=localhost port=5432 user=grpc password=password dbname=grpc sslmode=disable"

    
    DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get DB instance: %v", err)
    }
    
    if err = sqlDB.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    if err := DB.AutoMigrate(&User{}, &Transaction{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    
    log.Println("Connected to PostgreSQL database")
}
