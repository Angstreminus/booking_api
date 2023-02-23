package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection error:", err)
	} else {
		fmt.Println("We connected to a databse")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Appartment{})
}
