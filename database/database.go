package database

import (
  "gorm.io/driver/sqlite" 
  "gorm.io/gorm"
)

db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})


func ConnectDB(){}
