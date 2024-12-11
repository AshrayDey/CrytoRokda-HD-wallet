package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load the environment file :%v", err)
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DSN could not be retrieved")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("DB initialised successfully")
}

func createTable() {

	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		hashed_Password TEXT,
	);
	`
	result := DB.Exec(userTable)
	if result.Error != nil {
		log.Fatalf("Something wrong with statement:%v", result.Error)
	}

}
