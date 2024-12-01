package db

import (
	"loco-assignment/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	client *gorm.DB
}

func (pc *PostgresClient) InitializeClient() {
	dsn := os.Getenv("POSTGRES_DSN")
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	client.AutoMigrate(&models.Transaction{})

	pc.client = client
}

func (pc *PostgresClient) GetClient() *gorm.DB {
	return pc.client
}

func (pc *PostgresClient) CloseClient() {
	sqlDB, _ := pc.client.DB()
	sqlDB.Close()
}
