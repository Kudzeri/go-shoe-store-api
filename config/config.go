package config

import (
	"github.com/Kudzeri/go-shoe-store-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=mydb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	err = MigrateDB()
	if err != nil {
		panic("failed to migrate database schemas")
	}
}

func MigrateDB() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		return err
	}

	return nil
}
