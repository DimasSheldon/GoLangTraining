package database

import (
	"finalproject/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	dbPort   = "5432"
	dbName   = "thefinalproject"
	db       *gorm.DB
	err      error
)

func StartDB() {
	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	defer fmt.Println("Successfully Connected to Database")

	db.Debug().AutoMigrate(model.User{}, model.Comment{}, model.Photo{}, model.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
