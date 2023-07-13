package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"coinConversion/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func DatabaseInit() (*gorm.DB, error) {
	var database *gorm.DB

	dbUrl := os.Getenv("DATABASE_URL")

	database, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to the database, error: %v", err)
		return nil, err
	}

	err = database.AutoMigrate(&model.Coins{})
	if err != nil {
		database.AutoMigrate(model.Coins{})
	}

	err = database.AutoMigrate(&model.Logs{})
	if err != nil {
		database.AutoMigrate(model.Logs{})
	}

	var count int64
	database.Model(&model.Coins{}).Count(&count)
	if count == 0 {

		coins := []model.Coins{
			{Name: "Real", Abbreviation: "BRL", Symbol: "R$", CreatedAt: time.Now()},
			{Name: "Dolar", Abbreviation: "USD", Symbol: "$", CreatedAt: time.Now()},
			{Name: "Euro", Abbreviation: "EUR", Symbol: "€", CreatedAt: time.Now()},
			{Name: "Bitcoin", Abbreviation: "BTC", Symbol: "₿", CreatedAt: time.Now()},
		}

		for _, coin := range coins {
			err = database.Create(&coin).Error
			if err != nil {
				setError := fmt.Sprintf("error to insert coins %v:", err)
				log.Fatal(setError)
			}
		}

	}

	log.Println("Database connection successful...")
	return database, nil
}
