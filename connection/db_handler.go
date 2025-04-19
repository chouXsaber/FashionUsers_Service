package connection

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Failed load .env")
	}

	
	database_name := os.Getenv("DATABASE")
	serverhost    := os.Getenv("SERVER")
	username      := os.Getenv("USER")
	password      := os.Getenv("PASSWORD")
	port	      := os.Getenv("PORT")
	sslmode	      := os.Getenv("SSLMODE")
	timezone      := os.Getenv("TIMEZONE")

	dsn := "host=" + serverhost + 
		   " user=" + username + 
		   " password=" + password + 
		   " dbname=" + database_name + 
		   " port=" + port + 
		   " sslmode=" + sslmode + 
		   " TimeZone=" + timezone

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil{
		log.Fatal("Error connect db",err)
	}

	log.Println("Connect database")
	DB = db

}