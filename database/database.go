package database

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxthizeau/api-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database \n%v", err.Error())
		os.Exit(2)
	}

	log.Info("Connected to the database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Info("Running Migrations...")

	// TODO : Add migrations

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
