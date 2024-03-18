package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "rahasia"
	dbname   = "db-rest-api"
)

func StartDb() (*gorm.DB, error) {

	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil

	//err := db.AutoMigrate(models.Order{}, models.Item{})
	//if err != nil {
	//	return
	//}
}
