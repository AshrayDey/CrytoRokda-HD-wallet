package db

import (
	"HD-Wallet/models"
	"log"
)

func migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Wallet{})
	if err != nil {
		log.Fatalf("Migration error:%v", err)
	}
	log.Println("Database migrated successfully")
}
