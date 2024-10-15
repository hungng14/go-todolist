package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"todolist/config"
	"todolist/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func InitDB() {
	host := config.Config("DB_HOST")
	p := config.Config("DB_PORT")
	username := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	name := config.Config("DB_NAME")

	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing port string to number")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
		os.Exit(2)
	}

	log.Println("Database connected!")

	log.Println("Running migrations")
	db.AutoMigrate(&models.Task{})

	DB = DbInstance{
		Db: db,
	}
}
