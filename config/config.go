package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = ConnectDB()

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load `.env` file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9191 TimeZone=Asia/Dushanbe",
		dbHost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL database")
	}

	return db
}

func DisconnectDB(db *gorm.DB) {
	_db, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}

	_db.Close()
}
