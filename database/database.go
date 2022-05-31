package database

import (
	"log"
	"os"

	"github.com/jeffronworks/fiber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=192.168.95.41 user=jeffron password=root12 dbname=GoCrud port=5432 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Add migrations
	db.AutoMigrate(&models.Order{}, &models.Product{}, &models.User{})

	Database = DbInstance{Db: db}
}
