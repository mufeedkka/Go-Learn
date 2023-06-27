package database

import (
	"log"
	"os"

	"github.com/mufeedkka/goecommerce/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Faild yo connect DB \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected db success")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}

}
