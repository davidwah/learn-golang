package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-api/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "rahasia"
	dbname   = "db-rest-api"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err := db.AutoMigrate(models.Order{}, models.OrderItem{})
	if err != nil {
		return
	}

}

func GetDB() *gorm.DB {
	return db
}
