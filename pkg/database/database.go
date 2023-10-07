package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	dsn := "host=dpg-ckf2ncenpffc73bu1ji0-a.oregon-postgres.render.com user=admin password=l6sBPSz8Puf5dBfmzwnsyUPyaK0U9oIw dbname=samamander port=5432"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect Database")
	}
}
