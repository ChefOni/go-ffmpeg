package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Instance *gorm.DB
}

type DBLogs struct {
	ID                   uint   `gorm:"primaryKey;autoIncrement"`
	Date                 string
	StorageBucketProvider string
	UserIP               string
	RegisteredUser       bool
}

var DBInstance *DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("app_database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&DBLogs{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Successfully connected to SQLite database")
	DBInstance = &DB{Instance: db}
}
