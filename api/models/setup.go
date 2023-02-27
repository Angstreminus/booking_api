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

func Dburl() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error occured loading env: ", err)
	}
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_PASSWORD"))
}

func ConnectDataBase() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(Dburl()), &gorm.Config{})
}

func InitDB() {
	var err error
	DB, err = ConnectDataBase()
	if err != nil {
		log.Fatal(err)
	}
}
